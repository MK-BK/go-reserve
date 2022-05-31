package main

import (
	"go-reserve/api"
	"go-reserve/db"
	"go-reserve/middle"
	"go-reserve/models"
	"go-reserve/server/audit"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	initDatabase()
	startServer()
}

func startServer() {
	e := gin.Default()

	e.Use(middle.NewAuthHandler(), middle.NewErrorHandler())
	api.InitRoute(e)

	if err := e.Run(":10086"); err != nil {
		log.Fatal(err)
	}
}

func initDatabase() {
	db := db.GetDB()
	if err := db.AutoMigrate(
		models.User{},
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
