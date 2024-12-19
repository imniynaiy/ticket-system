package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imniynaiy/ticket-system/internal/errors"
	"github.com/imniynaiy/ticket-system/internal/model"
	"github.com/imniynaiy/ticket-system/internal/service"
)

func AddRouteRoutes(rg *gin.RouterGroup) {
	routes := rg.Group("/routes")
	routes.GET("", listRoutes)
}

func listRoutes(ctx *gin.Context) {
	var req model.UserListRoutesReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(errors.ErrInvalidRequest.HTTPStatus, model.NewErrorResponse(errors.ErrInvalidRequest))
		return
	}

	result, err := service.UserListRoutes(&req)
	if err != nil {
		ctx.JSON(errors.ErrInternalServerError.HTTPStatus, model.NewErrorResponse(errors.ErrInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, model.NewSuccessResponse(result))

}
