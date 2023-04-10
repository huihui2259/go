package entity

type User struct {
	ID         int    `gorm:"column:id;primaryKey" json:"id"`
	Phone      string `gorm:"column:phone" json:"phone"`
	NickName   string `gorm:"column:nick_name" json:"nickname"`
	Password   string `gorm:"column:password" json:"password"`
	Icon       string `gorm:"column:icon" json:"icon"`
	CreateTime string `gorm:"column:create_time" json:"createTime"`
	UpdateTime string `gorm:"column:update_time" json:"updateTime"`
}

func (User) TableName() string {
	return "tb_user" // 数据库表的名称
}
