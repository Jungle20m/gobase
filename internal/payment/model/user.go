package model

import "time"

type User struct {
	ID                     int        `gorm:"column:id"`
	UserName               string     `gorm:"column:user_name"`
	PhoneNumber            string     `gorm:"column:phone_number"`
	Email                  string     `gorm:"column:email"`
	Password               string     `gorm:"column:password"`
	PasswordSalt           string     `gorm:"column:password_salt"`
	PasswordHashAlgorithms string     `gorm:"column:password_hash_algorithms"`
	UserStatus             string     `gorm:"column:user_status"`
	RegistrationTime       *time.Time `gorm:"column:registration_time"`
	CreateTime             *time.Time `gorm:"column:create_time; autoCreateTime"`
	UpdateTime             *time.Time `gorm:"column:update_time; autoUpdateTime"`
}

func (User) TableName() string {
	return "user_account"
}
