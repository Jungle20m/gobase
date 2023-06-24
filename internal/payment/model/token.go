package model

import "time"

type Token struct {
	ID           int        `gorm:"column:id"`
	UserID       int        `gorm:"column:user_id"`
	ClientID     string     `gorm:"column:client_id"`
	IDToken      string     `gorm:"column:id_token"`
	AccessToken  string     `gorm:"column:access_token"`
	RefreshToken string     `gorm:"column:refresh_token"`
	CreateTime   *time.Time `gorm:"column:create_time; autoCreateTime"`
	UpdateTime   *time.Time `gorm:"column:update_time; autoUpdateTime"`
}

func (Token) TableName() string {
	return "user_token"
}
