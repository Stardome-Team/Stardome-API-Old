package errormiddleware

import (
	"github.com/gin-gonic/gin"
)

// ErrorHandlerMiddleware :
func ErrorHandlerMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			c.AbortWithStatusJSON(c.Writer.Status(), c.Errors)
		}
	}
}
