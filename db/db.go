package db

import (
	"github.com/towelong/vgo/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conn() *gorm.DB {
	var err error
	DB, err = gorm.Open(mysql.Open(global.Config.Data.Database.Source))
	if err != nil {
		panic(err)
	}
	return DB
}
