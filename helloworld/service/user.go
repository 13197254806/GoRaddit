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

func UserSignUp(paramSignIn *models.ParamSignUp) (err error) {
	err = mysql.IsUserNameExisted(paramSignIn.Username)
	if err != nil {
		return
	}

	err = mysql.InsertUser(map[string]interface{}{
		"UserId":   snowflake.GenerateID(),
		"Username": paramSignIn.Username,
		"Password": encryptPassword(paramSignIn.Password),
	})
	return
}
