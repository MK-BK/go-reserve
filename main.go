package main

import (
	"go-reserve/api"
	"go-reserve/middle"
	"go-reserve/models"
	"go-reserve/server/audit"
	"go-reserve/util/conf"
	DB "go-reserve/util/db"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	err := conf.LoadConf()
	if err != nil {
		panic(err)
	}

	initDatabase()

	startServer()
}

func startServer() {
	log.Print("Start go_reserve")

	gin.SetMode(gin.ReleaseMode)
	e := gin.Default()

	e.Use(middle.NewAuthHandler(), middle.NewErrorHandler())
	api.InitRoute(e)

	if err := e.Run(":" + conf.GetListenPort()); err != nil {
		panic(err)
	}
}

func initDatabase() {
	DB.InitDatabase()
	db := DB.GetDB()
	if err := db.AutoMigrate(
		models.Account{},
		models.Job{},
		models.Shop{},
		models.Session{},
		models.Message{},
		models.Commodity{},
		models.Request{},
		audit.AuditLog{},
	); err != nil {
		panic(err)
	}
}
