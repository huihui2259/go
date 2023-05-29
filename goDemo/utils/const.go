package utils

import (
	"time"
)

const (
	UserIDPrefix           = "user::"
	ShopIDPrefix           = "shop::"
	LockPrefix             = "lock::"
	VoucherPrefix          = "voucher::"
	SecKillVoucherPrefix   = "seckillvoucher::"
	SecKillOrderLockPrefix = "order::lock::"
	VoucherStockPrefix     = "stock::id::"
	SecKillUserIDPrefix    = "seckill::user::"
	BlogLikedIDPrefix      = "blog::liked::"
	FollowIDPrefix         = "follow::id::"
	SignPrefix             = "sign::"

	GetMysqlError     = "获取数据库错误"
	GetEmptyDataError = "获取到空数据"
	InsertMysqlError  = "插入数据库错误"
	DeleteMysqlError  = "删除数据库错误"
	UpdateMysqlError  = "更新数据库错误"
	GetRedisError     = "获取redis错误"
	SetRedisError     = "设置redis错误"
	DelRedisError     = "删除redis错误"
	ValueNullError    = "数据为空"

	// 秒杀活动中出现的错误
	StockNotEnough    = "库存不足"
	SecKillNotBegin   = "秒杀未开始"
	SecKillAlreadyEnd = "秒杀已结束"
	OrderCreateError  = "订单创建失败"
	RepeatOrder       = "重复下单"

	TimeFormat        = "2006-01-02 15:04:05"
	InitTime          = "2000-01-01 00:00:00"
	TimeFormat1       = "20060102"
	Local             = "Local"
	Year              = "2006"
	Month             = "01"
	Day               = "02"
	Hour              = "15"
	Minute            = "04"
	Second            = "05"
	OneDay            = 24 * 60 * 60
	CommonExpireTime  = 24 * time.Hour
	LockExpireTime    = 10 * time.Second // 为了表现出锁，故意将时间调大
	SecKillExpireTime = 24 * time.Hour

	CommonValue = "common value"

	// 查询一页数据默认值
	PageSize = 2

	// 一些常识值
	ZERO  = 0
	FALSE = 0
	TRUE  = 1
)
