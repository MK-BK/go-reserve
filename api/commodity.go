package api

import (
	"go-reserve/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func listCommoditys(c *gin.Context) {
	cs, err := GE.CommodityManager.List()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, cs)
}

func createCommodity(c *gin.Context) {
	var commdity models.Commodity
	if err := c.ShouldBind(&commdity); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := GE.CommodityManager.Create(&commdity); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

func getCommodity(c *gin.Context) {
	id := c.Param("id")
	commdity, err := GE.CommodityManager.Get(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, commdity)
}

func updateCommodity(c *gin.Context) {
	id := c.Param("id")
	var commdity models.Commodity
	if err := c.ShouldBind(&commdity); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := GE.CommodityManager.Update(id, &commdity)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

func deleteCommodity(c *gin.Context) {
	id := c.Param("id")
	if err := GE.CommodityManager.Delete(id); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
