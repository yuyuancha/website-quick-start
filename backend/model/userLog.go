package model

import "time"

// UserLog 會員紀錄
type UserLog struct {
	Id        int       `gorm:"primaryKey" json:"id"`
	UserId    int       `json:"userId"`
	User      User      `gorm:"foreignKey:UserId"`
	Operation int       `json:"operation"`
	IpAddress string    `json:"ipAddress"`
	Note      *string   `json:"note"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (*UserLog) TableName() string {
	return "user_logs"
}
