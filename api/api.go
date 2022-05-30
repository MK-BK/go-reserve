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
		users.POST("/", CreateUser)
		users.GET(":id", GetUser)
		users.PUT(":id", UpdateUser)
		users.DELETE(":id", DeleteUser)
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
		product := e.Group("/product")
		product.GET("/", ListProduct)
		product.POST("/", CreateProduct)
		product.GET(":id", GetProduct)
		product.PUT(":id", UpdateProduct)
		product.DELETE(":id", DeleteProduct)
	}

	{
		session := e.Group("/sessions")
		session.POST("/", CreateSession)
		session.GET("/", ListSessions)
		session.GET("/:id", GetSession)
		session.DELETE("/:id", DeleteSession)
		session.POST("/message", SendMessage)
	}
}

func jsonWithResult(c *gin.Context, value interface{}) {
}
