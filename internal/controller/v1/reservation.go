package v1

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
	reservations.POST("", createReservation)
}

func getReservation(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	userId := ctx.GetUint("UserId")
	if err != nil {
		ctx.JSON(errors.ErrInvalidRequest.HTTPStatus, model.NewErrorResponse(errors.ErrInvalidRequest))
		return
	}

	reservation, err := service.GetUserReservationWithDetails(uint(id), userId)
	if err != nil {
		ctx.JSON(errors.ErrNotFound.HTTPStatus, model.NewErrorResponse(errors.ErrNotFound))
		return
	}

	ctx.JSON(http.StatusOK, model.NewSuccessResponse(reservation))
}

func listReservations(ctx *gin.Context) {
	userId := ctx.GetUint("UserId")
	var req model.ListReservationsReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(errors.ErrInvalidRequest.HTTPStatus, model.NewErrorResponse(errors.ErrInvalidRequest))
		return
	}

	result, err := service.ListUserReservations(userId, &req)
	if err != nil {
		ctx.JSON(errors.ErrInternalServerError.HTTPStatus, model.NewErrorResponse(errors.ErrInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, model.NewSuccessResponse(result))
}

func createReservation(ctx *gin.Context) {
	userId := ctx.GetUint("UserId")
	var req model.CreateReservationReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(errors.ErrInvalidRequest.HTTPStatus, model.NewErrorResponse(errors.ErrInvalidRequest))
		return
	}

	result, err := service.CreateReservation(userId, &req)
	if err != nil {
		ctx.JSON(errors.ErrInternalServerError.HTTPStatus, model.NewErrorResponse(errors.ErrInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, model.NewSuccessResponse(result))
}
