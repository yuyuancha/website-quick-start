package server

import (
	"github.com/gin-gonic/gin"
	"github.com/yuyuancha/website-quick-start/controller"
)

func setRouter(router *gin.Engine) {
	router.GET("/ping", controller.CommonController.Ping)
}
