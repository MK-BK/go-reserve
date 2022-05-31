package api

import (
	"go-reserve/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListRequest(c *gin.Context) {
	rs, err := GE.RequestManager.List()
	if err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
	c.JSON(http.StatusOK, rs)
}

func GetRequest(c *gin.Context) {
	id := c.Param("id")
	request, err := GE.RequestManager.Get(id)
	if err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
	c.JSON(http.StatusOK, request)
}

func UpdateRequest(c *gin.Context) {
	id := c.Param("id")
	var requet models.Request
	if err := c.ShouldBind(&requet); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
	if err := GE.RequestManager.Update(id, &requet); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
}
