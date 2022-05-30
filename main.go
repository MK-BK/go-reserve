package main

import (
	"go-reserve/api"
	"go-reserve/common/db"
	"go-reserve/middle"
	"go-reserve/models"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	initDatabase()
	startServer()
}

func startServer() {
	e := gin.Default()
	e.Use(middle.NewErrorHandler())
	api.InitRoute(e)

	if err := e.Run(":10086"); err != nil {
		log.Fatal(err)
	}
}

func initDatabase() {
	db := db.GetDB()
	if err := db.AutoMigrate(models.User{}, models.Job{}, models.Product{}, models.Session{}, models.Message{}); err != nil {
		panic(err)
	}
}
