package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imniynaiy/ticket-system/internal/errors"
	"github.com/imniynaiy/ticket-system/internal/model"
	"github.com/imniynaiy/ticket-system/internal/service"
)

func AddSeatRoutes(rg *gin.RouterGroup) {
	seats := rg.Group("/seats")
	seats.POST("", createSeat)
	seats.GET("", listSeats)
	seats.GET(":id", getSeat)
	seats.PUT("", updateSeat)
	seats.DELETE(":id", deleteSeat)
}

func createSeat(ctx *gin.Context) {
	var req model.CreateSeatReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(errors.ErrInvalidRequest.HTTPStatus, model.NewErrorResponse(errors.ErrInvalidRequest))
		return
	}

	seat, err := service.CreateSeat(&req)
	if err != nil {
		ctx.JSON(errors.ErrInternalServerError.HTTPStatus, model.NewErrorResponse(errors.ErrInternalServerError))
		return
	}

	ctx.JSON(http.StatusCreated, model.NewSuccessResponse(seat))
}

func getSeat(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(errors.ErrInvalidRequest.HTTPStatus, model.NewErrorResponse(errors.ErrInvalidRequest))
		return
	}

	seat, err := service.GetSeat(uint(id))
	if err != nil {
		ctx.JSON(errors.ErrNotFound.HTTPStatus, model.NewErrorResponse(errors.ErrNotFound))
		return
	}

	ctx.JSON(http.StatusOK, model.NewSuccessResponse(seat))
}

func updateSeat(ctx *gin.Context) {
	var req model.UpdateSeatReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(errors.ErrInvalidRequest.HTTPStatus, model.NewErrorResponse(errors.ErrInvalidRequest))
		return
	}

	seat, err := service.UpdateSeat(&req)
	if err != nil {
		ctx.JSON(errors.ErrInternalServerError.HTTPStatus, model.NewErrorResponse(errors.ErrInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, model.NewSuccessResponse(seat))
}

func deleteSeat(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(errors.ErrInvalidRequest.HTTPStatus, model.NewErrorResponse(errors.ErrInvalidRequest))
		return
	}

	if err := service.DeleteSeat(uint(id)); err != nil {
		ctx.JSON(errors.ErrInternalServerError.HTTPStatus, model.NewErrorResponse(errors.ErrInternalServerError))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func listSeats(ctx *gin.Context) {
	var req model.ListSeatsReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(errors.ErrInvalidRequest.HTTPStatus, model.NewErrorResponse(errors.ErrInvalidRequest))
		return
	}

	result, err := service.ListSeats(&req)
	if err != nil {
		ctx.JSON(errors.ErrInternalServerError.HTTPStatus, model.NewErrorResponse(errors.ErrInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, model.NewSuccessResponse(result))
}
