package entity

type Shop struct {
	ID         int    `gorm:"column:id;primaryKey" json:"id"`
	Name       string `gorm:"column:name" json:"name"`
	Type       int    `gorm:"column:type_id" json:"type"`
	Images     string `gorm:"column:images" json:"images"`
	Area       string `gorm:"column:area" json:"area"`
	Adderss    string `gorm:"column:address" json:"address"`
	AvgPrice   int    `gorm:"column:avg_price" json:"avgPrice"`
	Sold       int    `gorm:"column:sold"  json:"sold"`
	Comments   int    `gorm:"column:comments" json:"comments"`
	Open       string `gorm:"column:open_hours" json:"open"`
	CreateTime string `gorm:"column:create_time" json:"createTime"`
	UpdateTime string `gorm:"column:update_time" json:"updateTime"`
}

func (Shop) TableName() string {
	return "tb_shop"
}
