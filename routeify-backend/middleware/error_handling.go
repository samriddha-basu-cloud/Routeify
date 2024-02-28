// middleware/error_handling.go
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandlerMiddleware(c *gin.Context) {
	c.Next()

	// Check for errors
	errors := c.Errors.ByType(gin.ErrorTypeAny)
	if len(errors) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}
}
