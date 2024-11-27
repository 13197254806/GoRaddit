package mysql

import (
	"errors"
)

type User struct {
	//gorm.Model
	ID uint `gorm:"primarykey"`

	//UserID int64
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
		err = errors.New("该用户已存在")
	} else if dbErr != nil {
		err = dbErr
	}
	return
}

func InsertUser(userInfo map[string]interface{}) (err error) {
	newUser := User{
		UserId:   userInfo["UserId"].(int64),
		Username: userInfo["Username"].(string),
		Password: userInfo["Password"].(string),
	}
	//fmt.Println("%v", newUser)
	return db.Create(&newUser).Error
}
