package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imniynaiy/ticket-system/internal/errors"
	"github.com/imniynaiy/ticket-system/internal/model"
	"github.com/imniynaiy/ticket-system/internal/service"
)

func AddSeatclassRoutes(rg *gin.RouterGroup) {
	seatclasses := rg.Group("/seatclasses")
	seatclasses.POST("", createSeatclass)
	seatclasses.GET("", listSeatclasses)
	seatclasses.GET(":id", getSeatclass)
	seatclasses.PUT("", updateSeatclass)
	seatclasses.DELETE(":id", deleteSeatclass)
}

func createSeatclass(ctx *gin.Context) {
	var req model.CreateSeatclassReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(errors.ErrInvalidRequest.HTTPStatus, model.NewErrorResponse(errors.ErrInvalidRequest))
		return
	}

	seatclass, err := service.CreateSeatclass(&req)
	if err != nil {
		ctx.JSON(errors.ErrInternalServerError.HTTPStatus, model.NewErrorResponse(errors.ErrInternalServerError))
		return
	}

	ctx.JSON(http.StatusCreated, model.NewSuccessResponse(seatclass))
}

func getSeatclass(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(errors.ErrInvalidRequest.HTTPStatus, model.NewErrorResponse(errors.ErrInvalidRequest))
		return
	}

	seatclass, err := service.GetSeatclass(uint(id))
	if err != nil {
		ctx.JSON(errors.ErrNotFound.HTTPStatus, model.NewErrorResponse(errors.ErrNotFound))
		return
	}

	ctx.JSON(http.StatusOK, model.NewSuccessResponse(seatclass))
}

func updateSeatclass(ctx *gin.Context) {
	var req model.UpdateSeatclassReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(errors.ErrInvalidRequest.HTTPStatus, model.NewErrorResponse(errors.ErrInvalidRequest))
		return
	}

	seatclass, err := service.UpdateSeatclass(&req)
	if err != nil {
		ctx.JSON(errors.ErrInternalServerError.HTTPStatus, model.NewErrorResponse(errors.ErrInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, model.NewSuccessResponse(seatclass))
}

func deleteSeatclass(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(errors.ErrInvalidRequest.HTTPStatus, model.NewErrorResponse(errors.ErrInvalidRequest))
		return
	}

	if err := service.DeleteSeatclass(uint(id)); err != nil {
		ctx.JSON(errors.ErrInternalServerError.HTTPStatus, model.NewErrorResponse(errors.ErrInternalServerError))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func listSeatclasses(ctx *gin.Context) {
	result, err := service.ListSeatclasses()
	if err != nil {
		ctx.JSON(errors.ErrInternalServerError.HTTPStatus, model.NewErrorResponse(errors.ErrInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, model.NewSuccessResponse(result))
}
