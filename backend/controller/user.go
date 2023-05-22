package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yuyuancha/website-quick-start/exception"
	"github.com/yuyuancha/website-quick-start/model"
	"github.com/yuyuancha/website-quick-start/service"
	"net/http"
)

var UserController User

// User 會員 Controller
type User struct{}

// Register 會員註冊 Api
func (User) Register(ctx *gin.Context) {
	var data model.UserRegisterData
	err := ctx.ShouldBind(&data)
	if err != nil {
		exception.SendException(ctx, "UNKNOWN")
		return
	}

	errorName := service.Register(ctx, data)
	if errorName != "" {
		exception.SendException(ctx, errorName)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "操作成功",
	})
}
