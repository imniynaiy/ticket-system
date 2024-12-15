package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seatclass, err := service.CreateSeatclass(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, seatclass)
}

func getSeatclass(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid seatclass ID"})
		return
	}

	seatclass, err := service.GetSeatclass(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, seatclass)
}

func updateSeatclass(ctx *gin.Context) {
	var req model.UpdateSeatclassReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	seatclass, err := service.UpdateSeatclass(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, seatclass)
}

func deleteSeatclass(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid seatclass ID"})
		return
	}

	if err := service.DeleteSeatclass(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func listSeatclasses(ctx *gin.Context) {
	result, err := service.ListSeatclasses()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
