package mysql

import (
	"errors"
	"github.com/jinzhu/gorm"
)

var (
	ErrorUserExist   = errors.New("用户已存在")
	ErrorInvalidUser = errors.New("用户名或密码错误")
	ErrorMysql       = errors.New("数据库错误")
)

type User struct {
	//gorm.Model
	ID uint `gorm:"primarykey"`

	UserId   int64
	Username string
	Password string
	Email    string
	Gender   int8
}

func (User) TableName() string {
	return "user"
}

func IsUserNameExisted(userName string) (err error) {
	var count int64
	dbErr := db.Model(&User{}).Where("username = ?", userName).Count(&count).Error
	if count > 0 {
		err = ErrorUserExist
	} else if dbErr != nil {
		err = ErrorMysql
	}
	return err
}

func IsUserExisted(user *User) (err error) {
	dbErr := db.Where("username = ? and password = ?", user.Username, user.Password).First(user).Error
	if errors.Is(dbErr, gorm.ErrRecordNotFound) {
		err = ErrorInvalidUser
	} else if dbErr != nil {
		err = ErrorMysql
	}
	return
}

func InsertUser(userInfo map[string]interface{}) (err error) {
	newUser := User{
		UserId:   userInfo["UserId"].(int64),
		Username: userInfo["Username"].(string),
		Password: userInfo["Password"].(string),
	}
	if dbErr := db.Create(&newUser).Error; dbErr != nil {
		err = ErrorMysql
	}
	return
}
