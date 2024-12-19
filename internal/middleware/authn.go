package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/imniynaiy/ticket-system/internal/errors"
	"github.com/imniynaiy/ticket-system/internal/model"
	"github.com/imniynaiy/ticket-system/internal/util"
)

const authHeader = "Authorization"

func Authenticationer() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		h := c.GetHeader(authHeader)
		_, token, found := strings.Cut(h, "Bearer ")
		if !found {
			c.AbortWithStatusJSON(errors.ErrInvalidToken.HTTPStatus, model.NewErrorResponse(errors.ErrInvalidToken))
			return
		}

		session, err := util.VerifyTokenWithRedis(token)
		if err != nil {
			c.AbortWithStatusJSON(errors.ErrInvalidToken.HTTPStatus, model.NewErrorResponse(errors.ErrInvalidToken))
			return
		}

		c.Set("UserId", uint(session.UserID))
		c.Set("IsAdmin", session.IsAdmin)

		c.Next()

		// after request
	}
}
