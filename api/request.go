package api

import (
	"go-reserve/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func listRequests(c *gin.Context) {
	rs, err := GE.RequestManager.List()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, rs)
}

func getRequest(c *gin.Context) {
	id := c.Param("id")
	request, err := GE.RequestManager.Get(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, request)
}

func updateRequest(c *gin.Context) {
	id := c.Param("id")
	var requet models.Request
	if err := c.ShouldBind(&requet); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if err := GE.RequestManager.Update(id, &requet); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
