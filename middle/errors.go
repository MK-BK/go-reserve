package middle

import (
	"github.com/gin-gonic/gin"
)

func NewErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			c.JSON(400, c.Errors[0])
		}
	}
}
