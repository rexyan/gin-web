package dao

import (
	"go.uber.org/zap"
	"web_app/models"
	"web_app/pkg/mysql"
)

type CommunityDao struct{}

func (com *CommunityDao) CommunityList() (community *[]models.Community, err error) {
	var communityInstance []models.Community
	sqlStr := "select community_id, community_name, introduction from community"
	if err = mysql.DB.Select(&communityInstance, sqlStr); err != nil {
		zap.L().Error("get user by name error", zap.Error(err))
		return nil, err
	} else {
		return &communityInstance, nil
	}
}
