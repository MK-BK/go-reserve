package middle

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			c.JSON(http.StatusInternalServerError, c.Errors.Last())
		}
	}
}
