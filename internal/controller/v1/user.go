package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imniynaiy/ticket-system/internal/log"
	"github.com/imniynaiy/ticket-system/internal/model"
	"github.com/imniynaiy/ticket-system/internal/service"
)

func AddUserRoutes(rg *gin.RouterGroup) {
	rg.POST("/login", userLoginController)
	rg.POST("/register", userRegisterController)
}

func userLoginController(c *gin.Context) {
	var user model.LoginReq
	err := c.ShouldBindBodyWithJSON(&user)
	if err != nil {
		log.Error("Failed to parse login user", log.String("err", err.Error()))
		c.Status(http.StatusBadRequest)
		return
	}
	token, err := service.Login(&user)
	if err != nil {
		log.Error("Failed to login", log.String("err", err.Error()))
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, &model.LoginResp{Token: token})
}

func userRegisterController(c *gin.Context) {
	var user model.RegisterReq
	err := c.ShouldBindBodyWithJSON(&user)
	if err != nil {
		log.Error("Failed to parse register user", log.String("err", err.Error()))
		c.Status(http.StatusBadRequest)
		return
	}
	err = service.Register(&user)
	if err != nil {
		log.Error("Failed to register", log.String("err", err.Error()))
		c.Status(http.StatusBadRequest)
		return
	}
	c.Status(http.StatusOK)
}
