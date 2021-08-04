package controller

import (
	"net/http"
	"strconv"
	"web_app/middleware"
	"web_app/pkg/response"
	v "web_app/pkg/validator"
	"web_app/service"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var CommunityService = new(service.CommunityService)

// CommunityListHandler 社区列表
// @Summary 获取社区列表接口
// @Description 获取社区列表接口
// @Tags 社区
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} []models.Community
// @Router /community [get]
func CommunityListHandler(c *gin.Context) {
	communityList, err := CommunityService.CommunityList()
	if err != nil {
		zap.L().Error("CommunityService CommunityList Error:", zap.Error(err))
		response.BuildResponse(c, nil, response.CommunityListError, 400)
		return
	}
	response.BuildSuccessResponse(c, communityList)
}

// CommunityDetailHandler 社区详情
// @Summary 获取社区详情接口
// @Description 获取社区详情接口
// @Tags 社区
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param id path string true "社区 ID"
// @Security ApiKeyAuth
// @Success 200 {object} models.Community
// @Router /community/{id} [get]
func CommunityDetailHandler(c *gin.Context) {
	communityIdStr := c.Param("id")
	communityID, err := strconv.ParseInt(communityIdStr, 10, 64)
	if err != nil {
		zap.L().Error("Query Param Community ID Error:", zap.Error(err))
		response.BuildResponse(c, nil, response.CommunityIDError, 400)
		return
	}
	community, err := CommunityService.CommunityByID(communityID)
	if err != nil {
		zap.L().Error("CommunityService CommunityByID Error:", zap.Error(err))
		response.BuildResponse(c, nil, response.CommunityDetailError, 400)
		return
	}
	response.BuildSuccessResponse(c, community)
}

// CreatePostHandler 创建帖子
// @Summary 创建帖子
// @Description 创建帖子
// @Tags 社区
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body validator.CreatePostValidator true "社区 ID"
// @Security ApiKeyAuth
// @Success 200 {object} models.Post
// @Router /post [post]
func CreatePostHandler(c *gin.Context) {
	// 参数校验
	createPostValidator := new(v.CreatePostValidator)
	err := c.ShouldBindJSON(createPostValidator)
	if err != nil {
		zap.L().Error("CreatePostValidator error", zap.Error(err))
		validationError, ok := err.(validator.ValidationErrors)
		if ok {
			response.BuildResponseWithMsg(c, nil, v.TransError(validationError), response.ParamError, http.StatusBadRequest)
			return
		}
		response.BuildResponse(c, nil, response.ParamError, http.StatusBadRequest)
		return
	}
	// 获取并转换 UserID
	ctxUserID, _ := c.Get(middleware.CtxUserID)
	userID, ok := ctxUserID.(int64)
	if !ok {
		zap.L().Error("UserID Conversion Int64 Error", zap.Error(err))
		response.BuildResponse(c, nil, response.CreatePostError, http.StatusBadRequest)
		return
	}
	// 创建帖子
	err = CommunityService.CreatePost(*createPostValidator, userID)
	if err != nil {
		zap.L().Error("CreatePost Error", zap.Error(err))
		response.BuildResponse(c, nil, response.CreatePostError, http.StatusBadRequest)
		return
	}
	response.BuildSuccessResponse(c, "ok")
}
