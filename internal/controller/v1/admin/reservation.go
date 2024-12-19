package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imniynaiy/ticket-system/internal/errors"
	"github.com/imniynaiy/ticket-system/internal/model"
	"github.com/imniynaiy/ticket-system/internal/service"
)

func AddReservationRoutes(rg *gin.RouterGroup) {
	reservations := rg.Group("/reservations")
	reservations.GET("", listReservations)
	reservations.GET(":id", getReservation)
}

func getReservation(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(errors.ErrInvalidRequest.HTTPStatus, model.NewErrorResponse(errors.ErrInvalidRequest))
		return
	}

	reservation, err := service.GetReservationWithDetails(uint(id))
	if err != nil {
		ctx.JSON(errors.ErrNotFound.HTTPStatus, model.NewErrorResponse(errors.ErrNotFound))
		return
	}

	ctx.JSON(http.StatusOK, model.NewSuccessResponse(reservation))
}

func listReservations(ctx *gin.Context) {
	var req model.ListReservationsReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(errors.ErrInvalidRequest.HTTPStatus, model.NewErrorResponse(errors.ErrInvalidRequest))
		return
	}

	result, err := service.ListReservations(&req)
	if err != nil {
		ctx.JSON(errors.ErrInternalServerError.HTTPStatus, model.NewErrorResponse(errors.ErrInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, model.NewSuccessResponse(result))
}
