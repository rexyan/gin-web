package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"web_app/pkg/response"
	"web_app/service"
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
