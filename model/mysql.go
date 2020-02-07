package model

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// MYSQL 是一个mysql连接实例
var MYSQL *gorm.DB

// MysqlInit 用来初始化连接
func MysqlInit(connstring string) {
	db, err := gorm.Open("mysql", connstring)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)

	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	//空闲
	db.DB().SetMaxIdleConns(20)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	MYSQL = db

	migration()

}
