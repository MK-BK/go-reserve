package api

import (
	"go-reserve/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListShop(c *gin.Context) {
	productes, err := GE.ProductManage.List()
	if err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}

	c.JSON(http.StatusOK, productes)
}

func CreateShop(c *gin.Context) {
	var product *models.Shop
	if err := c.ShouldBind(&product); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}

	if err := GE.ProductManage.Create(product); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
}

func GetShop(c *gin.Context) {
	id := c.Param("id")
	proudct, err := GE.ProductManage.Get(id)
	if err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
	c.JSON(http.StatusOK, proudct)
}

func UpdateShop(c *gin.Context) {
	id := c.Param("id")
	var product *models.Shop
	if err := GE.ProductManage.Update(id, product); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}

}

func DeleteShop(c *gin.Context) {
	id := c.Param("id")
	if err := GE.ProductManage.Delete(id); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
}
