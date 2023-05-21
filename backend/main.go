package main

import (
	"fmt"

	"github.com/yuyuancha/website-quick-start/config"
	"github.com/yuyuancha/website-quick-start/model"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println(config.App)
	fmt.Println(config.Driver.Mysql)
	var user model.User
	user.Username = "uuu123"
	user.Nickname = "芋圓茶2"
	user.Password = "abc123456"
	user.Create()
}
