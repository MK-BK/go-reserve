package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	loadConf()
	initDatabase()
}

func initDatabase() {
	var err error

	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "123456", "127.0.0.1:6666", "go_reserve"),
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func loadConf() {

}

func GetDB() *gorm.DB {
	return db
}
