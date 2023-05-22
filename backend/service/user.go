package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yuyuancha/website-quick-start/config"
	"github.com/yuyuancha/website-quick-start/model"
)

// Register 註冊會員
func Register(ctx *gin.Context, data model.UserRegisterData) string {
	var user = model.User{
		Username: data.Username,
		Nickname: data.Nickname,
		Password: data.Password,
		Email:    data.Email,
	}

	if errorName := validateUserRegisterData(user); errorName != "" {
		return errorName
	}

	if err := user.Create(); err != nil {
		return "UNKNOWN"
	}

	var userLog = model.UserLog{
		User:      user,
		Operation: config.UserLogOperation["LOGIN"],
		IpAddress: ctx.ClientIP(),
	}

	_ = userLog.Create()

	return ""
}

// validateUserRegisterData 驗證會員註冊資料
func validateUserRegisterData(user model.User) string {
	isExist, err := user.IsExist()
	if err != nil {
		return "UNKNOWN"
	} else if isExist {
		return "USERNAME_EXIST"
	}

	isUsed, nicknameErr := user.IsNicknameUsed()
	if nicknameErr != nil {
		return "UNKNOWN"
	} else if isUsed {
		return "NICKNAME_USED"
	}

	return ""
}
