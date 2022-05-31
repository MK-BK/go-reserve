package api

import (
	"go-reserve/server"

	"github.com/gin-gonic/gin"
)

var GE = server.Env

func InitRoute(e *gin.Engine) {
	{
		users := e.Group("/users")
		users.POST("/login", UserLogin)
		users.GET("/", ListUsers)
		users.POST("/", CreateUser)
		users.GET(":id", GetUser)
		users.PUT(":id", UpdateUser)
		users.DELETE(":id", DeleteUser)
	}

	{
		request := e.Group("/request")
		request.GET("/", ListRequest)
		request.GET("/:id", GetRequest)
		request.PUT("/:id", UpdateRequest)
	}

	{
		job := e.Group("/jobs")
		job.GET("/", ListJob)
		job.POST("/", CreateJob)
		job.GET(":id", GetJob)
		job.PUT(":id", UpdateJob)
		job.DELETE(":id", DeleteJob)
	}

	{
		shop := e.Group("/shops")
		shop.GET("/", ListShop)
		shop.POST("/", CreateShop)
		shop.GET(":id", GetShop)
		shop.PUT(":id", UpdateShop)
		shop.DELETE(":id", DeleteShop)
	}

	{
		commodity := e.Group("/commodity")
		commodity.GET("/", ListCommodity)
		commodity.POST("/", CreateCommodity)
		commodity.GET(":id", GetCommodity)
		commodity.PUT(":id", UpdateCommodity)
		commodity.DELETE(":id", DeleteCommodity)
	}

	{
		session := e.Group("/sessions")
		session.POST("/", CreateSession)
		session.GET("/", ListSessions)
		session.GET("/:id", GetSession)
		session.DELETE("/:id", DeleteSession)
		session.POST("/message", SendMessage)
	}

	e.GET("/audit_log", ListAuditLog)
}

func jsonWithResult(c *gin.Context, value interface{}) {
}
