package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	// gorm 进行mysql连接 返回DB+error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	// 自动根据结构体创建表
	//err = DB.AutoMigrate(&UserInfo{}, &Video{}, &Comment{}, &UserLogin{})
	//if err != nil {
	//	log.Panicln("err:", err.Error())
	//}
}
