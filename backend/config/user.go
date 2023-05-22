package config

// UserLogOperation 會員紀錄操作碼
var UserLogOperation map[string]int

func init() {
	UserLogOperation = make(map[string]int)

	UserLogOperation["LOGIN"] = 1  // 會員登入
	UserLogOperation["LOGOUT"] = 2 // 會員登出
}
