package service

import (
	"web_app/dao"
	"web_app/models"
)

type CommunityService struct{}

var CommunityDao = new(dao.CommunityDao)

func (com *CommunityService) CommunityList() (community *[]models.Community, err error) {
	return CommunityDao.CommunityList()
}
