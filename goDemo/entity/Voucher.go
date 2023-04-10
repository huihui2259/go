package entity

type Voucher struct {
	ID          int    `gorm:"column:id;primaryKey" json:"id"`
	ShopID      int    `gorm:"column:shop_id" json:"shopId"`
	Title       string `gorm:"column:title" json:"title"`
	SubTitle    string `gorm:"column:sub_title" json:"subTitle"`
	Rules       string `gorm:"column:rules" json:"rules"`
	PayValue    int    `gorm:"column:pay_value" json:"payValue"`
	ActualValue int    `gorm:"column:actual_value" json:"actualValue"`
	Type        int    `gorm:"column:type" json:"type"`
	Status      int    `gorm:"column:status" json:"status"`
	CreateTime  string `gorm:"column:create_time" json:"createTime"`
	UpdateTime  string `gorm:"column:update_time" json:"updateTime"`
	Stock       int    `gorm:"-" json:"stock"`
	BeginTime   string `gorm:"-" json:"beginTime"`
	EndTime     string `gorm:"-" json:"endTime"`
}

func (Voucher) TableName() string {
	return "tb_voucher"
}

type SecKillVoucher struct {
	VoucherID  int    `gorm:"column:voucher_id;primaryKey"  json:"id"`
	Stock      int    `gorm:"column:stock"  json:"stock"`
	CreateTime string `gorm:"column:create_time" json:"createTime"`
	BeginTime  string `gorm:"column:begin_time" json:"beginTime"`
	EndTime    string `gorm:"column:end_time" json:"endTime"`
	UpdateTime string `gorm:"column:update_time" json:"updateTime"`
}

func (SecKillVoucher) TableName() string {
	return "tb_seckill_voucher"
}

type VoucherOrder struct {
	ID         int    `gorm:"column:id"`
	UserID     int    `gorm:"column:user_id"`
	VoucherID  int    `gorm:"column:voucher_id"`
	PayType    int    `gorm:"column:pay_type"`
	Status     int    `gorm:"column:status"`
	CreateTime string `gorm:"column:create_time"`
	PayTime    string `gorm:"column:pay_time"`
	UseTime    string `gorm:"column:use_time"`
	RefundTime string `gorm:"column:refund_time"`
	UpdateTime string `gorm:"column:update_time"`
}

func (VoucherOrder) TableName() string {
	return "tb_voucher_order"
}
