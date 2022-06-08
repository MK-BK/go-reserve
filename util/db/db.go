package db

import (
	"fmt"
	"go-reserve/util/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	InitDatabase()
}

func InitDatabase() {
	var err error

	if err = conf.LoadConf(); err != nil {
		panic(err)
	}

	c := conf.GetMysqlSection()
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.DBName),
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

func GetDB() *gorm.DB {
	return db
}
