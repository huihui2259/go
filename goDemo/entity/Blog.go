package entity

type Blog struct {
	ID         int    `gorm:"column:id" json:"id"`
	ShopID     int    `gorm:"column:shop_id" json:"shopId"`
	UserID     int    `gorm:"column:user_id" json:"userId"`
	Title      string `gorm:"column:title" json:"title"`
	Images     string `gorm:"column:images" json:"images"`
	Content    string `gorm:"column:content" json:"content"`
	Liked      int    `gorm:"column:liked" json:"liked"`
	Comments   int    `gorm:"column:comments" json:"comments"`
	CreateTime string `gorm:"column:create_time" json:"createTime"`
	UpdateTime string `gorm:"column:update_time" json:"updateTime"`
	Like       bool   `gorm:"-"`
}

func (Blog) TableName() string {
	return "tb_blog"
}

type Follow struct {
	ID           int    `gorm:"column:id"`
	UserID       int    `gorm:"column:user_id"`
	FollowUserID int    `gorm:"column:follow_user_id"`
	CreateTime   string `gorm:"column:create_time"`
}

func (Follow) TableName() string {
	return "tb_follow"
}
