package service

import (
	"crypto/md5"
	"encoding/hex"
	"web_app/dao"
	"web_app/middleware"
	"web_app/models"
	"web_app/pkg/jwt"
	"web_app/pkg/snowflake"
	"web_app/pkg/validator"
	"web_app/settings"

	"github.com/gin-gonic/gin"
)

type UserService struct{}

var userDao = new(dao.UserDao)

// 根据用户名判断用户是否存在
func (uc *UserService) UserExistByName(username string) bool {
	if user, err := userDao.UserByName(username); user != nil || err != nil {
		return true
	}
	return false
}

// 注册用户
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

// 密码加密
func (uc *UserService) EncryptPassword(password string) string {
	hash := md5.New()
	hash.Write([]byte(settings.Config.ServerConfig.Secret))
	sum := hash.Sum([]byte(password))
	return hex.EncodeToString(sum)
}

// 用户登录
func (uc *UserService) LoginService(param *validator.LoginValidator) (user *models.User, err error) {
	user, err = userDao.Login(param.Username, uc.EncryptPassword(param.Password))
	if err != nil {
		return nil, err
	}
	return user, nil
}

// 生成 JWT Token
func (uc *UserService) GenerateJwtToken(userID int64, userName string) (accessToken, refreshToken string) {
	accessToken, _ = jwt.GenerateJwtAccessToken(userID, userName)
	refreshToken, _ = jwt.GenerateJwtRefreshToken(userID, userName)
	return
}

// 生成用户 ID
func (uc *UserService) GetUserByID(userID int64) *models.User {
	if user, err := userDao.UserByID(userID); err != nil {
		return nil
	} else {
		return user
	}
}

// 获取当前登录的用户
func (uc *UserService) GetCurrentUser(c *gin.Context) *models.User {
	// 从请求上下文中获取 user_id
	userID, _ := c.Get(middleware.CtxUserID)
	intUserID := userID.(int64)
	return uc.GetUserByID(intUserID)
}
