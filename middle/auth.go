package middle

import (
	"github.com/gin-gonic/gin"
)

func NewAuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// token := c.GetHeader("token")
		// if token == "" {
		// 	c.AbortWithError(http.StatusUnauthorized, nil)
		// }

		// c.Next()
		// if len(c.Errors) > 0 {
		// 	c.JSON(400, c.Errors[0])
		// }
	}
}
