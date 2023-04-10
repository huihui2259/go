package mysql

import (
	"fmt"
	"goDemo/conf"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

// 初始化db
func init() {
	cfg := conf.Config
	log.Println(cfg)
	dbParams := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=False&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.DBName)
	var err error
	Db, err = gorm.Open("mysql", dbParams)
	if err != nil {
		log.Fatal("连接数据库错误", err)
	}
	Db.DB().SetMaxOpenConns(100)
	fmt.Println("database init on port ", cfg.Host)
}
