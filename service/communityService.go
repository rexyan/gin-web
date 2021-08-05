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

// 获取社区列表
func (com *CommunityService) CommunityList() (community *[]models.Community, err error) {
	return CommunityDao.CommunityList()
}

// 根据 ID 查询社区
func (com *CommunityService) CommunityByID(communityID int64) (community *models.Community, err error) {
	return CommunityDao.CommunityByID(communityID)
}

// 创建帖子
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

// 根据 ID 查询帖子
func (com *CommunityService) PostByID(postID int64) (post *models.Post, err error) {
	return CommunityDao.PostByID(postID)
}

// 根据 ID 查询帖子详情
func (com *CommunityService) PostDetail(id int64) (postDetail *validator.PostDetail, err error) {
	// 查询 post
	post, err := com.PostByID(id)
	if err != nil {
		return nil, err
	}

	// 查询 user
	userDao := new(dao.UserDao)
	user, err := userDao.UserByID(post.UserId)
	if err != nil {
		return nil, err
	}

	// 查询 community
	communityDao := new(dao.CommunityDao)
	community, err := communityDao.CommunityByID(post.CommunityId)
	if err != nil {
		return nil, err
	}

	// 组装数据
	postDetail = &validator.PostDetail{
		PostID:        post.PostID,
		Title:         post.Title,
		Content:       post.Content,
		UserId:        post.UserId,
		UserName:      user.UserName,
		CommunityId:   post.CommunityId,
		CommunityName: community.CommunityName,
		CreateTime:    post.CreateTime,
	}
	return postDetail, nil
}

// 分页查询帖子列表
func (com *CommunityService) PostByPage(page, pageSize int64) (postList *[]validator.PostDetail, err error) {
	userDao := new(dao.UserDao)
	communityDao := new(dao.CommunityDao)
	postListData := make([]validator.PostDetail, 0, pageSize)

	postByPage, err := CommunityDao.PostByPage(page, pageSize)
	if err != nil {
		return nil, err
	}

	for _, post := range postByPage {
		// 查询用户
		user, err := userDao.UserByID(post.UserId)
		if err != nil {
			continue
		}
		// 查询社区
		community, err := communityDao.CommunityByID(post.CommunityId)
		if err != nil {
			continue
		}
		// 组装数据
		postDetail := validator.PostDetail{
			PostID:        post.PostID,
			Title:         post.Title,
			Content:       post.Content,
			UserId:        post.UserId,
			UserName:      user.UserName,
			CommunityId:   post.CommunityId,
			CommunityName: community.CommunityName,
			CreateTime:    post.CreateTime,
		}
		// 数据添加到 slice 中
		postListData = append(postListData, postDetail)

	}
	return &postListData, nil
}
