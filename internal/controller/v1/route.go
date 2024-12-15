package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imniynaiy/ticket-system/internal/model"
	"github.com/imniynaiy/ticket-system/internal/service"
)

func AddRouteRoutes(rg *gin.RouterGroup) {
	routes := rg.Group("/routes")
	routes.POST("", createRoute)
	routes.GET("", listRoutes)
	routes.GET(":id", getRoute)
	routes.PUT("", updateRoute)
	routes.DELETE(":id", deleteRoute)
}

// CreateRoute godoc
// @Summary Create a new route
// @Description Create a new route with the provided details
// @Tags routes
// @Accept json
// @Produce json
// @Param route body model.CreateRouteReq true "Route details"
// @Success 201 {object} model.Route
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /routes [post]
func createRoute(ctx *gin.Context) {
	var req model.CreateRouteReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	route, err := service.CreateRoute(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, route)
}

// GetRoute godoc
// @Summary Get a route by ID
// @Description Get route details by route ID
// @Tags routes
// @Produce json
// @Param id path int true "Route ID"
// @Success 200 {object} model.Route
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /routes/{id} [get]
func getRoute(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid route ID"})
		return
	}

	route, err := service.GetRoute(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, route)
}

// UpdateRoute godoc
// @Summary Update a route
// @Description Update a route with the provided details
// @Tags routes
// @Accept json
// @Produce json
// @Param id path int true "Route ID"
// @Param route body model.UpdateRouteReq true "Route details"
// @Success 200 {object} model.Route
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /routes/{id} [put]
func updateRoute(ctx *gin.Context) {
	var req model.UpdateRouteReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	route, err := service.UpdateRoute(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, route)
}

// DeleteRoute godoc
// @Summary Delete a route
// @Description Delete a route by ID
// @Tags routes
// @Produce json
// @Param id path int true "Route ID"
// @Success 204 "No Content"
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /routes/{id} [delete]
func deleteRoute(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid route ID"})
		return
	}

	if err := service.DeleteRoute(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// ListRoutes godoc
// @Summary List all routes with pagination
// @Description Get a paginated list of all routes
// @Tags routes
// @Produce json
// @Param page query int false "Page number (default: 1)" minimum(1)
// @Param page_size query int false "Items per page (default: 10)" minimum(1) maximum(100)
// @Success 200 {object} model.ListRoutesResp
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /routes [get]
func listRoutes(ctx *gin.Context) {
	var req model.ListRoutesReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set default values if not provided
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10 // Default page size
	}

	result, err := service.ListRoutes(req.Page, req.PageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
