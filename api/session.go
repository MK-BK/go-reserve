package api

import (
	"go-reserve/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateSession(c *gin.Context) {
	var session models.Session
	if err := GE.SessionManage.CreateSession(&session); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
}

func ListSessions(c *gin.Context) {
	sessions, err := GE.SessionManage.ListSession()
	if err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
	c.JSON(http.StatusOK, sessions)
}

func GetSession(c *gin.Context) {
	var id string
	session, err := GE.SessionManage.GetSession(id)
	if err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
	c.JSON(http.StatusOK, session)
}

func DeleteSession(c *gin.Context) {
	var id string
	if err := GE.SessionManage.DeleteSession(id); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
}

func SendMessage(c *gin.Context) {
	var message models.Message
	if err := GE.SessionManage.SendMessage(&message); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
}
