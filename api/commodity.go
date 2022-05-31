package api

import (
	"go-reserve/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListCommodity(c *gin.Context) {
	cs, err := GE.CommodityManage.List()
	if err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
	c.JSON(http.StatusOK, cs)
}

func CreateCommodity(c *gin.Context) {
	var commdity models.Commodity
	if err := c.ShouldBind(&commdity); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}

	if err := GE.CommodityManage.Create(&commdity); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
}

func GetCommodity(c *gin.Context) {
	id := c.Param("id")
	commdity, err := GE.CommodityManage.Get(id)
	if err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
	c.JSON(http.StatusOK, commdity)
}

func UpdateCommodity(c *gin.Context) {
	id := c.Param("id")
	var commdity models.Commodity
	if err := c.ShouldBind(&commdity); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}

	err := GE.CommodityManage.Update(id, &commdity)
	if err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
}

func DeleteCommodity(c *gin.Context) {
	id := c.Param("id")
	if err := GE.CommodityManage.Delete(id); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
}
