package model

import (
	"fmt"
)

func GetUser(name string) *User {
	var ret User
	err := Conn.First(&ret, "name = ?", name).Error
	if err != nil {
		fmt.Println("err:", err.Error())
	}
	fmt.Println("User retrieved from database:", ret) // 添加这行以查看检索到的用户信息
	return &ret
}
