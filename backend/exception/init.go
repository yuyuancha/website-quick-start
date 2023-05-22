package exception

import (
	"github.com/gin-gonic/gin"
	"github.com/yuyuancha/website-quick-start/config"
)

// Base 基本例外訊息內容
type Base struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// SendException 傳送錯誤
func SendException(ctx *gin.Context, errorName string) {
	if errorName == "" {
		return
	}

	switch errorName {
	default:
		sendCommonException(ctx, errorName)
	}
}

// sendCommonException 傳送通用錯誤
func sendCommonException(ctx *gin.Context, errorName string) {
	exception := Base{
		Code:    config.ExceptionCode[errorName],
		Message: config.ExceptionMessage[errorName],
	}

	ctx.JSON(config.ExceptionHttpCode[errorName], exception)
}
