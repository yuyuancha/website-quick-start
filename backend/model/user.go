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

// IsExist 檢查會員是否存在
func (u *User) IsExist() bool {
	var count int64
	driver.MySql.Table(u.TableName()).Where("username", u.Username).Count(&count)

	return count != 0
}

// IsNicknameUsed 檢查暱稱是否使用過
func (u *User) IsNicknameUsed() bool {
	var count int64
	driver.MySql.Table(u.TableName()).Where("nickname", u.Nickname).Count(&count)

	return count != 0
}

// Create 建立會員資料
func (u *User) Create() error {
	hash, err := util.GeneratePasswordHash(u.Password)
	if err != nil {
		return errors.New("密碼加密發生錯誤")
	}

	u.Password = hash

	err = driver.MySql.Create(u).Error
	if err != nil {
		return err
	}

	return nil
}
