package base

import (
	"fmt"
	"gin-cli/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// 数据库初始化设置

var DB *gorm.DB

func InitDataBase() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
			util.GetConfig("db", "username"),
			util.GetConfig("db", "password"),
			util.GetConfig("db", "host"),
			util.GetConfig("db", "port"),
			util.GetConfig("db", "name"),
			util.GetConfig("db","charSet")),
		DefaultStringSize: 256,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err) // 打印输出错误信息
	}
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}