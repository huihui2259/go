package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type AdminUser struct {
	ID       int    `gorm:"column:id"`
	UserName string `gorm:"column:user_name"`
	Password string `gorm:"column:password"`
	Tel      string `gorm:"column:tel"`
	Addr     string `gorm:"column:addr"`
}

func (AdminUser) TableName() string {
	return "admin_user" // 数据库表的名称
}

func main() {
	fmt.Println("开始连接")
	db, err := gorm.Open("mysql", "root:123456@tcp(localhost:3306)/user?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	if err != nil {
		fmt.Println("连接失败")
		fmt.Println(err)
	}
	u1 := &AdminUser{1, "xiaoming", "123", "123456", "beijing"}
	u2 := &AdminUser{2, "xiaohong", "123", "12345", "shanghai"}
	db.Create(u1)
	db.Create(u2)

	fmt.Println("连接成功！！！")
}
