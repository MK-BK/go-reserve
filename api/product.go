package api

import (
	"go-reserve/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListProduct(c *gin.Context) {
	productes, err := GE.ProductManage.List()
	if err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}

	c.JSON(http.StatusOK, productes)
}

func CreateProduct(c *gin.Context) {
	var product *models.Product
	if err := GE.ProductManage.Create(product); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	proudct, err := GE.ProductManage.Get(id)
	if err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
	c.JSON(http.StatusOK, proudct)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product *models.Product
	if err := GE.ProductManage.Update(id, product); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}

}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := GE.ProductManage.Delete(id); err != nil {
		c.Errors = append(c.Errors, &gin.Error{Err: err})
		return
	}
}
