package api

import (
	"go-reserve/server/audit"
	"net/http"

	"github.com/gin-gonic/gin"
)

func listAuditLog(c *gin.Context) {
	logs, err := audit.ListLogs()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, logs)
}
