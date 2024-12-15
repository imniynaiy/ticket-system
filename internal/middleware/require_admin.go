package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !c.GetBool("IsAdmin") {
			c.Status(http.StatusForbidden)
			c.Abort()
		}
	}
}
