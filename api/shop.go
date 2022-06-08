package api

import (
	"go-reserve/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func listShops(c *gin.Context) {
	productes, err := GE.ProductManager.List()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, productes)
}

func createShop(c *gin.Context) {
	var product *models.Shop
	if err := c.ShouldBind(&product); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := GE.ProductManager.Create(product); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

func getShop(c *gin.Context) {
	id := c.Param("id")
	proudct, err := GE.ProductManager.Get(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, proudct)
}

func updateShop(c *gin.Context) {
	id := c.Param("id")
	var product *models.Shop
	if err := GE.ProductManager.Update(id, product); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

}

func deleteShop(c *gin.Context) {
	id := c.Param("id")
	if err := GE.ProductManager.Delete(id); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
