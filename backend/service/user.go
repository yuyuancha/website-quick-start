package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/yuyuancha/website-quick-start/config"
	"github.com/yuyuancha/website-quick-start/model"
)

// Register 註冊會員
func Register(ctx *gin.Context, data model.UserRegisterData) error {
	var user = model.User{
		Username: data.Username,
		Nickname: data.Nickname,
		Password: data.Password,
		Email:    data.Email,
	}

	if err := validateUserRegisterData(user); err != nil {
		return err
	}

	if err := user.Create(); err != nil {
		return err
	}

	var userLog = model.UserLog{
		User:      user,
		Operation: config.UserLogOperation["LOGIN"],
		IpAddress: ctx.ClientIP(),
	}

	_ = userLog.Create()

	return nil
}

// validateUserRegisterData 驗證會員註冊資料
func validateUserRegisterData(user model.User) error {
	if user.IsExist() {
		return errors.New("會員帳號已存在")
	}

	if user.IsNicknameUsed() {
		return errors.New("會員暱稱已被使用")
	}

	return nil
}
