package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/imniynaiy/ticket-system/internal/errors"
	"github.com/imniynaiy/ticket-system/internal/model"
)

func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !c.GetBool("IsAdmin") {
			c.AbortWithStatusJSON(errors.ErrForbidden.HTTPStatus, model.NewErrorResponse(errors.ErrForbidden))
			return
		}
		c.Next()
	}
}
