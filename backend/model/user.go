package model

import (
	"errors"
	"time"

	"github.com/yuyuancha/website-quick-start/driver"
	"github.com/yuyuancha/website-quick-start/util"
)

// User 會員結構
type User struct {
	Id        int       `gorm:"primaryKey" json:"id"`
	Username  string    `json:"username"`
	Nickname  string    `json:"nickname"`
	Password  string    `json:"password"`
	Email     *string   `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (*User) TableName() string {
	return "users"
}

// Create 建立會員資料
func (u *User) Create() (*User, error) {
	hash, err := util.GeneratePasswordHash(u.Password)
	if err != nil {
		return nil, errors.New("密碼加密發生錯誤")
	}

	u.Password = hash

	err = driver.MySql.Create(u).Error
	if err != nil {
		return nil, err
	}

	return u, nil
}
