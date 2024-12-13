package service

import (
	"crypto/md5"
	"encoding/hex"
	"test.com/helloworld/dao/mysql"
	"test.com/helloworld/models"
	"test.com/helloworld/pkgs/snowflake"
)

const secret = "songchangtian"

func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}

func UserSignUp(paramSignUp *models.ParamSignUp) (err error) {
	err = mysql.IsUserNameExisted(paramSignUp.Username)
	if err != nil {
		return
	}

	err = mysql.InsertUser(map[string]interface{}{
		"UserId":   snowflake.GenerateID(),
		"Username": paramSignUp.Username,
		"Password": encryptPassword(paramSignUp.Password),
	})
	return
}

func UserSignIn(paramSignIn *models.ParamSignIn) (err error) {
	err = mysql.IsUserExisted(
		paramSignIn.Username,
		encryptPassword(paramSignIn.Password))
	return err
}
