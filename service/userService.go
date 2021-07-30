package service

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"web_app/dao"
	"web_app/middleware"
	"web_app/models"
	"web_app/pkg/jwt"
	"web_app/pkg/snowflake"
	"web_app/pkg/validator"
	"web_app/settings"
)

type UserService struct{}

var userDao = new(dao.UserDao)

func (uc *UserService) UserExistByName(username string) bool {
	// 判断用户是否存在
	if user, err := userDao.GetByName(username); user != nil || err != nil {
		return true
	}
	return false
}

func (uc *UserService) RegisterService(param *validator.RegisterValidator) (err error) {
	// 构建实例
	user := &models.User{
		UserID:   snowflake.GenID(),
		UserName: param.Username,
		Password: uc.EncryptPassword(param.Password),
		Email:    param.Email,
		Gender:   param.Gender,
	}
	// 保存
	return userDao.Insert(user)
}

func (uc *UserService) EncryptPassword(password string) string {
	hash := md5.New()
	hash.Write([]byte(settings.Config.ServerConfig.Secret))
	sum := hash.Sum([]byte(password))
	return hex.EncodeToString(sum)
}

func (uc *UserService) LoginService(param *validator.LoginValidator) (user *models.User, err error) {
	user, err = userDao.Login(param.Username, uc.EncryptPassword(param.Password))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserService) GenerateJwtToken(userID int64, userName string) (accessToken, refreshToken string) {
	accessToken, _ = jwt.GenerateJwtAccessToken(userID, userName)
	refreshToken, _ = jwt.GenerateJwtRefreshToken(userID, userName)
	return
}

func (uc *UserService) GetUserByID(userID int64) *models.User {
	if user, err := userDao.GetByID(userID); err != nil {
		return nil
	} else {
		return user
	}
}

func (uc *UserService) GetCurrentUser(c *gin.Context) *models.User {
	// 从请求上下文中获取 user_id
	userID, _ := c.Get(middleware.CtxUserID)
	intUserID := userID.(int64)
	return uc.GetUserByID(intUserID)
}
