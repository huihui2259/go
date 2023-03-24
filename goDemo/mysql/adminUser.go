package mysql

import (
	"log"
)

// 对admin_user表的操作
type User struct {
	ID         int    `gorm:"column:id"`
	Phone      string `gorm:"column:phone"`
	NickName   string `gorm:"column:nick_name"`
	Password   string `gorm:"column:password"`
	Icon       string `gorm:"column:icon"`
	CreateTime string `gorm:"column:create_time"`
	UpdateTime string `gorm:"column:update_time"`
}

func (User) TableName() string {
	return "tb_user" // 数据库表的名称
}

func SaveUser(user *User) error {
	log.Println(user)
	result := Db.Create(user)
	return result.Error
}

func GetUser(id int) (*User, error) {
	user := &User{}
	result := Db.Find(user, "id = ?", id)
	log.Println(user)
	return user, result.Error
}

func Delete(id int) error {

	result := Db.Where("id = ?", id).Delete(&User{})
	return result.Error
}

func Update(user *User) error {
	result := Db.Model(user).Where("id = ?", user.ID).Update("update_time", user.UpdateTime)
	return result.Error
}
