package entity

type Shop struct {
	ID         int    `gorm:"column:id"`
	Name       string `gorm:"column:name"`
	Type       int    `gorm:"column:type_id"`
	Images     string `gorm:"column:images"`
	Area       string `gorm:"column:area"`
	Adderss    string `gorm:"column:address"`
	AvgPrice   int    `gorm:"column:avg_price"`
	Sold       int    `gorm:"column:sold"`
	Comments   int    `gorm:"column:comments"`
	Open       string `gorm:"column:open_hours"`
	CreateTime string `gorm:"column:create_time"`
	UpdateTime string `gorm:"column:update_time"`
}

func (Shop) TableName() string {
	return "tb_shop"
}
