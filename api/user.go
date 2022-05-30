package api

import (
	"go-reserve/models"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	u := &models.User{}
	if err := GE.UserManger.Login(u); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}

	if err := GE.UserManger.Create(&user); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	if _, err := GE.UserManger.Get(id); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	u := &models.User{}
	if err := GE.UserManger.Update(id, u); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := GE.UserManger.Delete(id); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
}
