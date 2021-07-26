package controller

import (
	"net/http"
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
			c.JSON(http.StatusBadRequest, validatorError.Translate(v.Trans))
			return
		}
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if service.UserExistByName(registerValidator.Username) {
		zap.L().Error("user exist")
		c.JSON(http.StatusBadRequest, "user exist")
		return
	}
	if err := service.RegisterService(registerValidator); err != nil {
		zap.L().Error("user register error", zap.Error(err))
		c.JSON(http.StatusBadRequest, "user register error")
		return
	}
	c.JSON(http.StatusOK, "ok")
}
