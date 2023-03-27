package entity

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
