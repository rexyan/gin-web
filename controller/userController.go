package controller

import (
	"net/http"
	"web_app/pkg/response"
	v "web_app/pkg/validator"
	"web_app/service"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	// 参数校验
	registerValidator := new(v.RegisterValidator)
	if err := c.ShouldBindJSON(&registerValidator); err != nil {
		zap.L().Error("registerValidator error", zap.Error(err))
		validatorError, ok := err.(validator.ValidationErrors)
		if ok {
			response.BuildResponseWithMsg(c, nil, validatorError.Translate(v.Trans), response.ParamError, http.StatusBadRequest)
			return
		}
		response.BuildResponse(c, nil, response.ParamError, http.StatusBadRequest)
		return
	}
	if service.UserExistByName(registerValidator.Username) {
		zap.L().Error("user exist")
		response.BuildResponse(c, nil, response.UserExist, http.StatusBadRequest)
		return
	}
	if err := service.RegisterService(registerValidator); err != nil {
		zap.L().Error("user register error", zap.Error(err))
		response.BuildResponse(c, nil, response.UserRegisterError, http.StatusBadRequest)
		return
	}
	response.BuildSuccessResponse(c, "ok")
}

func LoginHandler(c *gin.Context) {
	loginValidator := new(v.LoginValidator)
	if err := c.ShouldBindJSON(loginValidator); err != nil {
		zap.L().Error("loginValidator error", zap.Error(err))
		validatorError, ok := err.(validator.ValidationErrors)
		if ok {
			response.BuildResponseWithMsg(c, nil, validatorError.Translate(v.Trans), response.ParamError, http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if user, err := service.LoginService(loginValidator); err != nil {
		zap.L().Error("user login error", zap.Error(err))
		response.BuildResponse(c, nil, response.UserLoginError, http.StatusBadRequest)
		return
	} else {
		response.BuildSuccessResponse(c, user)
	}
}
