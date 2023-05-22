package config

import "net/http"

var ExceptionHttpCode = make(map[string]int)
var ExceptionCode = make(map[string]int)
var ExceptionMessage = make(map[string]string)

func init() {
	ExceptionHttpCode["UNKNOWN"] = http.StatusInternalServerError
	ExceptionHttpCode["USERNAME_EXIST"] = http.StatusBadRequest
	ExceptionHttpCode["NICKNAME_USED"] = http.StatusBadRequest

	ExceptionCode["UNKNOWN"] = 19999
	ExceptionCode["USERNAME_EXIST"] = 11000
	ExceptionCode["NICKNAME_USED"] = 11001

	ExceptionMessage["UNKNOWN"] = "未知的錯誤，請聯繫管理員"
	ExceptionMessage["USERNAME_EXIST"] = "會員帳號已存在"
	ExceptionMessage["NICKNAME_USED"] = "會員暱稱已被使用"
}
