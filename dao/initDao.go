package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB

func Init() {
	//数据库连接初始化
	var err error
	dsn := "root:ZXzx580315@tcp(127.0.0.1:3306)/douyin_data?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("err:", err.Error())
	} else {
		fmt.Println("数据库连接成功")
	}

	sqlDB, _ := Db.DB()
	sqlDB.SetMaxOpenConns(100)
	//设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)
	//连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
}
