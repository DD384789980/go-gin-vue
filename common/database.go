package common

import (
	"go-gin-vue/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

// 开启连接池
func InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:3K7HozUxGp@(192.168.10.231:3306)/gin_vue?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("fail to connect database, err:" + err.Error())
	}
	
	// gorm 自动创建数据表
	db.AutoMigrate(&model.User{})

	DB = db
	return db
}

// 定义方法获取DB实例
func GetDB() *gorm.DB {
	return DB
}