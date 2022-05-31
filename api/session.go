package api

import (
	"go-reserve/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListSessions(c *gin.Context) {
	sessions, err := GE.SessionManage.List()
	if err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
	c.JSON(http.StatusOK, sessions)
}

func CreateSession(c *gin.Context) {
	var session models.Session
	if err := c.ShouldBind(&session); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}

	if err := GE.SessionManage.Create(&session); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
}

func GetSession(c *gin.Context) {
	id := c.Param("id")
	session, err := GE.SessionManage.Get(id)
	if err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
	c.JSON(http.StatusOK, session)
}

func DeleteSession(c *gin.Context) {
	var id string
	if err := GE.SessionManage.Delete(id); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
}

func SendMessage(c *gin.Context) {
	var message models.Message
	if err := c.ShouldBind(&message); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}

	if err := GE.SessionManage.SendMessage(&message); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
}
