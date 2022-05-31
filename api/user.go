package api

import (
	"go-reserve/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListUsers(c *gin.Context) {
	users, err := GE.UserManger.List()
	if err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
	c.JSON(http.StatusOK, users)
}

func UserLogin(c *gin.Context) {
	u := &models.User{}
	token, err := GE.UserManger.Login(u)
	if err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
	c.JSON(http.StatusOK, token)
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
	user, err := GE.UserManger.Get(id)
	if err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
	c.JSON(http.StatusOK, user)
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
