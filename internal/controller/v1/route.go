package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := service.UserListRoutes(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)

}
