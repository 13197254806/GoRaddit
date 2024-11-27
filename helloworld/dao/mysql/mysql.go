package mysql

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"go.uber.org/zap"
	"test.com/helloworld/settings"
)

var db *gorm.DB

func Init(mysqlConfig *settings.MysqlConfig) (err error) {
	//dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		mysqlConfig.UserName,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.DBName)
	db, err = gorm.Open("mysql", dsn)
	//err = db.DB().Ping()
	return
}

func CloseDB() {
	if err := db.Close(); err != nil {
		zap.L().DPanic("failed to close database", zap.Error(err))
	}
	return
}
