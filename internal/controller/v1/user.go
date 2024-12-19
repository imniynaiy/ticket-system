package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/imniynaiy/ticket-system/internal/errors"
	"github.com/imniynaiy/ticket-system/internal/log"
	"github.com/imniynaiy/ticket-system/internal/model"
	"github.com/imniynaiy/ticket-system/internal/service"
)

var validate = validator.New()

func AddUserRoutes(rg *gin.RouterGroup) {
	rg.POST("/login", userLoginController)
	rg.POST("/register", userRegisterController)
}

func userLoginController(c *gin.Context) {
	var user model.LoginReq
	err := c.ShouldBindBodyWithJSON(&user)
	if err != nil {
		log.Error("Failed to parse login user", log.String("err", err.Error()))
		c.JSON(errors.ErrInvalidRequest.HTTPStatus, model.NewErrorResponse(errors.ErrInvalidRequest))
		return
	}
	if err := validate.Struct(&user); err != nil {
		log.Error("Failed to validate login user", log.String("err", err.Error()))
		c.JSON(errors.ErrInvalidRequest.HTTPStatus, model.NewErrorResponse(errors.ErrInvalidRequest))
		return
	}
	token, err := service.Login(&user)
	if err != nil {
		log.Error("Failed to login", log.String("err", err.Error()))
		if appErr, ok := err.(*errors.AppError); ok {
			c.JSON(appErr.HTTPStatus, model.NewErrorResponse(appErr))
			return
		}
		c.JSON(errors.ErrInternalServerError.HTTPStatus, model.NewErrorResponse(errors.ErrInternalServerError))
		return
	}
	c.JSON(http.StatusOK, model.NewSuccessResponse(&model.LoginResp{Token: token}))
}

func userRegisterController(c *gin.Context) {
	var user model.RegisterReq
	err := c.ShouldBindBodyWithJSON(&user)
	if err != nil {
		log.Error("Failed to parse register user", log.String("err", err.Error()))
		c.JSON(errors.ErrInvalidRequest.HTTPStatus, model.NewErrorResponse(errors.ErrInvalidRequest))
		return
	}
	if err := validate.Struct(&user); err != nil {
		log.Error("Failed to validate register user", log.String("err", err.Error()))
		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Email" {
				c.JSON(errors.ErrInvalidEmail.HTTPStatus, model.NewErrorResponse(errors.ErrInvalidEmail))
				return
			}
			if err.Field() == "Password" {
				c.JSON(errors.ErrWeakPassword.HTTPStatus, model.NewErrorResponse(errors.ErrWeakPassword))
				return
			}
			if err.Field() == "Username" {
				c.JSON(errors.ErrInvalidUsername.HTTPStatus, model.NewErrorResponse(errors.ErrInvalidUsername))
				return
			}
		}
		c.JSON(errors.ErrInvalidRequest.HTTPStatus, model.NewErrorResponse(errors.ErrInvalidRequest))
		return
	}
	err = service.Register(&user)
	if err != nil {
		log.Error("Failed to register", log.String("err", err.Error()))
		if appErr, ok := err.(*errors.AppError); ok {
			c.JSON(appErr.HTTPStatus, model.NewErrorResponse(appErr))
			return
		}
		c.JSON(errors.ErrInternalServerError.HTTPStatus, model.NewErrorResponse(errors.ErrInternalServerError))
		return
	}
	c.JSON(http.StatusOK, model.NewSuccessResponse(nil))
}
