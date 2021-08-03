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

var userService = new(service.UserService)

// RefreshTokenHandler 获取 RefreshToken
// @Summary 获取 RefreshToken
// @Description 获取 RefreshToken
// @Tags 认证
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} string
// @Router /refreshToken [get]
func RefreshTokenHandler(c *gin.Context) {
	// 无需再次验证 jwt token，已经通过中间件验证
	//refreshToken := c.Request.Header.Get(middleware.AuthorizationKey)
	//if _, err := jwt.ParseToken(refreshToken); err != nil {
	//	response.BuildResponse(c, nil, response.RefreshTokenError, http.StatusBadRequest)
	//	return
	//}
	user := userService.GetCurrentUser(c)
	accessToken, _ := userService.GenerateJwtToken(user.UserID, user.UserName)
	response.BuildSuccessResponse(c, accessToken)
}

// RegisterHandler 用户注册
// @Summary 用户注册
// @Description 用户注册
// @Tags 认证
// @Accept application/json
// @Produce application/json
// @Param object body v.RegisterValidator true "注册参数"
// @Security ApiKeyAuth
// @Success 200 {object} string
// @Router /register [post]
func RegisterHandler(c *gin.Context) {
	// 参数校验
	registerValidator := new(v.RegisterValidator)
	if err := c.ShouldBindJSON(&registerValidator); err != nil {
		zap.L().Error("registerValidator error", zap.Error(err))
		validatorError, ok := err.(validator.ValidationErrors)
		if ok {
			response.BuildResponseWithMsg(c, nil, v.TransError(validatorError), response.ParamError, http.StatusBadRequest)
			return
		}
		response.BuildResponse(c, nil, response.ParamError, http.StatusBadRequest)
		return
	}
	if userService.UserExistByName(registerValidator.Username) {
		zap.L().Error("user exist")
		response.BuildResponse(c, nil, response.UserExist, http.StatusBadRequest)
		return
	}
	if err := userService.RegisterService(registerValidator); err != nil {
		zap.L().Error("user register error", zap.Error(err))
		response.BuildResponse(c, nil, response.UserRegisterError, http.StatusBadRequest)
		return
	}
	response.BuildSuccessResponse(c, "ok")
}

// LoginHandler 登录
// @Summary 登录
// @Description 登录
// @Tags 认证
// @Accept application/json
// @Produce application/json
// @Param object body v.LoginValidator true "登录信息"
// @Security ApiKeyAuth
// @Success 200 {object} string
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	loginValidator := new(v.LoginValidator)
	if err := c.ShouldBindJSON(loginValidator); err != nil {
		zap.L().Error("loginValidator error", zap.Error(err))
		validatorError, ok := err.(validator.ValidationErrors)
		if ok {
			response.BuildResponseWithMsg(c, nil, v.TransError(validatorError), response.ParamError, http.StatusBadRequest)
			return
		}
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if user, err := userService.LoginService(loginValidator); err != nil {
		zap.L().Error("user login error", zap.Error(err))
		response.BuildResponse(c, nil, response.UserLoginError, http.StatusBadRequest)
		return
	} else {
		accessToken, refreshToken := userService.GenerateJwtToken(user.UserID, user.UserName)
		response.BuildSuccessResponse(c, gin.H{
			"user":          user,
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		})
	}
}
