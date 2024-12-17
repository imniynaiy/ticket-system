package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/imniynaiy/ticket-system/internal/util"
)

const authHeader = "Authorization"

func Authenticationer() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		h := c.GetHeader(authHeader)
		_, token, found := strings.Cut(h, "Bearer ")
		if !found {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		session, err := util.VerifyTokenWithRedis(token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("UserId", uint(session.UserID))
		c.Set("IsAdmin", session.IsAdmin)

		c.Next()

		// after request
	}
}
