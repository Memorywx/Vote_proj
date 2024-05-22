package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 数据库操作

var Conn *gorm.DB

func MysqlConnect() {
	my := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "wangxiang666.", "127.0.0.1:3306", "vote")
	conn, err := gorm.Open(mysql.Open(my), &gorm.Config{})
	if err != nil {
		fmt.Println("err:", err)
		panic(err)
	}
	Conn = conn
}

func Close() {
	db, _ := Conn.DB()
	_ = db.Close()
}
