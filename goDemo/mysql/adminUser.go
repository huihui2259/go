package mysql

import (
	"log"
)

// 对admin_user表的操作
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

func SaveUser(user *AdminUser) error {
	log.Println(user)
	result := Db.Create(user)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetUser(id int) (*AdminUser, error) {
	user := &AdminUser{}
	result := Db.Find(user, "id = ?", id)
	if result.Error != nil {
		return user, result.Error
	}
	log.Println(user)
	return user, nil
}

func Delete(id int) error {

	result := Db.Where("id = ?", id).Delete(&AdminUser{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func Update(user *AdminUser) error {
	result := Db.Model(user).Where("id = ?", user.ID).Update("addr", user.Addr)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
