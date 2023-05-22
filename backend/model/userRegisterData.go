package model

// UserRegisterData 會員註冊資料
type UserRegisterData struct {
	Username string  `json:"username"`
	Nickname string  `json:"nickname"`
	Password string  `json:"password"`
	Email    *string `json:"email"`
}
