package service

import (
	"test.com/helloworld/dao/mysql"
	"test.com/helloworld/models"
	"test.com/helloworld/pkgs/snowflake"
)

func UserSignUp(paramSignIn *models.ParamSignIn) (err error) {
	err = mysql.IsUserNameExisted(paramSignIn.Username)
	if err != nil {
		return
	}
	err = mysql.InsertUser(map[string]interface{}{
		"UserID":   snowflake.GenerateID(),
		"Username": paramSignIn.Username,
		"Password": paramSignIn.Password,
	})
	return
}
