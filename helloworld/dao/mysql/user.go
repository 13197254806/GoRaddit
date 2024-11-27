package mysql

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID   int64 `gorm:"column: user_id"`
	Username string
	Password string
	Email    string
	Gender   string
}

func (User) TableName() string {
	return "user"
}

func IsUserNameExisted(userName string) (err error) {
	return db.Model(&User{}).Where("username = ?", userName).Error
}

func InsertUser(userInfo map[string]interface{}) (err error) {
	//return db.Model(&User{}).Create(userInfo).Error
	newUser := User{
		UserID:   userInfo["userID"].(int64),
		Username: userInfo["Username"].(string),
		Password: userInfo["Password"].(string),
	}
	return db.Create(&newUser).Error
}
