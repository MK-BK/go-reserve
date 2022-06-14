package api

import (
	"go-reserve/server"

	"github.com/gin-gonic/gin"
)

var GE = server.Env

func InitRoute(e *gin.Engine) {
	e.POST("/login", login)
	e.POST("/register", register)

	e.GET("/accounts", listAccounts)
	e.GET("/accounts/:id", getAccount)
	e.PUT("/accounts/:id", updateAccount)
	e.DELETE("/accounts/:id", deleteAccount)

	e.GET("/requests", listRequests)
	e.GET("/requests/:id", getRequest)
	e.PUT("/requests/:id", updateRequest)

	e.GET("/jobs", listJobs)
	e.POST("job", createJob)
	e.GET("/jobs/:id", getJob)
	e.PUT("/jobs/:id", updateJob)
	e.DELETE("/jobs/:id", deleteJob)

	e.GET("/shops", listShops)
	e.POST("/shop", createShop)
	e.GET("/shops/:id", getShop)
	e.PUT("/shops/:id", updateShop)
	e.DELETE("/shops/:id", deleteShop)

	e.GET("/commodity", listCommoditys)
	e.POST("/commodity", createCommodity)
	e.GET("/commodity/:id", getCommodity)
	e.PUT("/commodity/:id", updateCommodity)
	e.DELETE("/commodity/:id", deleteCommodity)

	e.GET("/sessions", listSessions)
	e.POST("/sessions", createSession)
	e.GET("/sessions/:id", getSession)
	e.DELETE("/sessions/:id", deleteSession)
	e.POST("/sessions/message", sendMessage)

	e.GET("/audit_log", listAuditLog)

	e.POST("/avatar/:resource/:id", uploadAvatar)
	e.GET("/avatar/:resource/:id", getAvatar)

	e.POST("/upload/", uploadFile)
}

func jsonWithResult(c *gin.Context, value interface{}) {
}
