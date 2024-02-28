// middleware/authorization.go
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminOnlyMiddleware(c *gin.Context) {
	role := c.GetString("role")
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Admin role required"})
		c.Abort()
		return
	}
	c.Next()
}
