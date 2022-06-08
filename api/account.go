package api

import (
	"go-reserve/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	var account models.Account
	if err := c.ShouldBind(&account); err != nil {
		c.Error(err)
		return
	}

	token, err := GE.AccountManager.Login(&account)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, token)
}

func register(c *gin.Context) {
	var account models.Account
	if err := c.ShouldBind(&account); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := GE.AccountManager.Create(&account); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

func listAccounts(c *gin.Context) {
	accounts, err := GE.AccountManager.List()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, accounts)
}

func getAccount(c *gin.Context) {
	id := c.Param("id")
	account, err := GE.AccountManager.Get(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, account)
}

func updateAccount(c *gin.Context) {
	id := c.Param("id")
	u := &models.Account{}
	if err := GE.AccountManager.Update(id, u); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

func deleteAccount(c *gin.Context) {
	id := c.Param("id")
	if err := GE.AccountManager.Delete(id); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
