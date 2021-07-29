package dao

import (
	"database/sql"
	"web_app/models"
	"web_app/pkg/mysql"

	"go.uber.org/zap"
)

type UserDao struct{}

func (dao *UserDao) Register(userInstance *models.User) (err error) {
	// 用户注册
	return dao.Insert(userInstance)
}

func (dao *UserDao) Insert(userInstance *models.User) (err error) {
	// 用户新增
	sqlStr := "INSERT INTO `gin_web`.`user`(`user_id`, `username`, `password`, `email`, `gender`) VALUES (?, ?, ?, ?, ?);"
	_, err = mysql.DB.Exec(sqlStr, userInstance.UserID, userInstance.UserName, userInstance.Password, userInstance.Email, userInstance.Gender)
	return err
}

func (dao *UserDao) GetByName(name string) (user *models.User, err error) {
	// 根据用户名查询用户
	var userInstance models.User
	sqlStr := "select user_id, username, password, email, gender from user where username=?"
	if err := mysql.DB.Get(&userInstance, sqlStr, name); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			zap.L().Error("get user by name error", zap.Error(err))
			return nil, err
		}
	}
	return &userInstance, nil
}

func (dao *UserDao) GetByID(userID int64) (user *models.User, err error) {
	// 根据用户ID查询用户
	var userInstance models.User
	sqlStr := "select user_id, username, password, email, gender from user where user_id=?"
	if err := mysql.DB.Get(&userInstance, sqlStr, userID); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			zap.L().Error("get user by id error", zap.Error(err))
			return nil, err
		}
	}
	return &userInstance, nil
}

func (dao *UserDao) Login(username, encryptPassword string) (user *models.User, err error) {
	// 用户登录
	var userInstance models.User
	sqlStr := "select user_id, username, password, email, gender from user where username=? and password=?"
	if err := mysql.DB.Get(&userInstance, sqlStr, username, encryptPassword); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Error("user login error", zap.Error(err))
			return nil, err
		}
	}
	return &userInstance, nil
}
