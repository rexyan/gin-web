package controller

import (
	"strconv"
	"web_app/pkg/response"
	"web_app/service"

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
