package conf

import (
	"os"
	"path/filepath"

	"github.com/go-ini/ini"
)

var config Config

type (
	Config struct {
		ListenPort string
		Mysql      `json:"mysql"`
	}

	Mysql struct {
		User     string
		Password string
		Host     string
		DBName   string
	}
)

func LoadConf() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	c, err := ini.Load(filepath.Join(dir, "util", "conf", "config.ini"))
	if err != nil {
		return err
	}

	section := c.Section("mysql")
	config.Mysql.Host = section.Key("host").String()
	config.Mysql.User = section.Key("user").String()
	config.Mysql.Password = section.Key("password").String()
	config.Mysql.DBName = section.Key("database").String()

	config.ListenPort = c.Section("").Key("ListenPort").String()
	return nil
}

func GetMysqlSection() Mysql {
	return config.Mysql
}

func GetListenPort() string {
	return config.ListenPort
}
