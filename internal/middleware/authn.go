package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/imniynaiy/ticket-system/internal/util"
)

var authHeader = "Authorization"

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

		c.AddParam("UserId", strconv.FormatUint(uint64(session.UserID), 10))
		c.AddParam("Role", strconv.Itoa(session.Role))

		c.Next()

		// after request
	}
}
