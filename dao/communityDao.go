package dao

import (
	"web_app/models"
	"web_app/pkg/mysql"

	"go.uber.org/zap"
)

type CommunityDao struct{}

func (com *CommunityDao) CommunityList() (community *[]models.Community, err error) {
	communityListInstance := new([]models.Community)
	sqlStr := "select community_id, community_name, introduction, create_time, update_time from community"
	if err = mysql.DB.Select(communityListInstance, sqlStr); err != nil {
		zap.L().Error("get community list error", zap.Error(err))
		return nil, err
	} else {
		return communityListInstance, nil
	}
}

func (com *CommunityDao) CommunityByID(communityID int64) (community *models.Community, err error) {
	communityInstance := new(models.Community)
	sqlStr := "select community_id, community_name, introduction, create_time, update_time from community where community_id = ?"
	if err := mysql.DB.Get(communityInstance, sqlStr, communityID); err != nil {
		zap.L().Error("get community by id error", zap.Error(err))
		return nil, err
	} else {
		return communityInstance, nil
	}
}

func (com *CommunityDao) CreatePost(newPost models.Post) (err error) {
	sqlStr := "INSERT INTO gin_web.post (post_id, title, content, user_id, community_id) VALUES (?,?,?,?,?);"
	if _, err := mysql.DB.Exec(sqlStr, newPost.PostID, newPost.Title, newPost.Content, newPost.UserId, newPost.CommunityId); err != nil {
		return err
	}
	return nil
}
