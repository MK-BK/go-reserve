package api

import (
	"go-reserve/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func listSessions(c *gin.Context) {
	sessions, err := GE.SessionManager.List()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, sessions)
}

func createSession(c *gin.Context) {
	var session models.Session
	if err := c.ShouldBind(&session); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := GE.SessionManager.Create(&session); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

func getSession(c *gin.Context) {
	id := c.Param("id")
	session, err := GE.SessionManager.Get(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, session)
}

func deleteSession(c *gin.Context) {
	var id string
	if err := GE.SessionManager.Delete(id); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

func sendMessage(c *gin.Context) {
	var message models.Message
	if err := c.ShouldBind(&message); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := GE.SessionManager.SendMessage(&message); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
