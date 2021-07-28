package service

import (
	"crypto/md5"
	"encoding/hex"
	"web_app/dao"
	"web_app/models"
	"web_app/pkg/snowflake"
	"web_app/pkg/validator"
	"web_app/settings"
)

var userDao = new(dao.UserDao)

func UserExistByName(username string) bool {
	// 判断用户是否存在
	if user, err := userDao.GetByName(username); user != nil || err != nil {
		return true
	}
	return false
}

func RegisterService(param *validator.RegisterValidator) (err error) {
	// 构建实例
	user := &models.User{
		UserID:   snowflake.GenID(),
		UserName: param.Username,
		Password: EncryptPassword(param.Password),
		Email:    param.Email,
		Gender:   param.Gender,
	}
	// 保存
	return userDao.Insert(user)
}

func EncryptPassword(password string) string {
	hash := md5.New()
	hash.Write([]byte(settings.Config.ServerConfig.Secret))
	sum := hash.Sum([]byte(password))
	return hex.EncodeToString(sum)
}

func LoginService(param *validator.LoginValidator) (user *models.User, err error) {
	user, err = userDao.Login(param.Username, EncryptPassword(param.Password))
	if err != nil {
		return nil, err
	}
	return user, nil
}
