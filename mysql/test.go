package mysql

import (
	"fmt"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Userbase struct { // 需要修改的字段映射
	UserId    uint32    `gorm:"column:USER_ID;PRIMARY_KEY"`
	UserName  string    `gorm:"column:USER_NAME"`
	LoginDate time.Time `gorm:"column:LOGIN_DATE"`
	ClientVar *string   `gorm:"column:CLIENT_VER"`
}

func (Userbase) TableName() string {
	return "userbase" // 数据库表的名称
}

type config struct {
	user   string
	pass   string
	adrr   string
	port   string
	dbname string
}

func sqlDb() *gorm.DB {
	conf := &config{
		user:   "root",    // 用户名
		pass:   "",        // 密码
		adrr:   localhost, // 地址
		port:   "3306",    // 端口
		dbname: "user",    // 数据库名称
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local", conf.user, conf.pass, conf.adrr, conf.port, conf.dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 打印所有sql
	})
	if err != nil {
		panic(err)
	}
	return db
}

func TestClientDb(t *testing.T) {
	db := sqlDb()
	ub := &Userbase{}
	err := db.Where("USER_ID = ?", 1187918).Find(&ub).Error
	if err != nil {
		panic(err)
	}
	fmt.Printf("userBase:%+v", ub)
}
