package service

import (
	"web_app/dao"
	"web_app/models"
	"web_app/pkg/snowflake"
	"web_app/pkg/validator"

	"go.uber.org/zap"
)

type CommunityService struct{}

var CommunityDao = new(dao.CommunityDao)

func (com *CommunityService) CommunityList() (community *[]models.Community, err error) {
	return CommunityDao.CommunityList()
}

func (com *CommunityService) CommunityByID(communityID int64) (community *models.Community, err error) {
	return CommunityDao.CommunityByID(communityID)
}

func (com *CommunityService) CreatePost(param validator.CreatePostValidator, userID int64) (err error) {
	// 验证 CommunityId
	if _, err := com.CommunityByID(int64(param.CommunityId)); err != nil {
		zap.L().Error("CreatePost Check CommunityID Error", zap.Error(err))
		return err
	}
	// 组装 Post
	newPost := models.Post{
		PostID:      snowflake.GenID(),
		UserId:      userID,
		Title:       param.Title,
		Content:     param.Title,
		CommunityId: param.CommunityId,
	}
	return CommunityDao.CreatePost(newPost)
}
