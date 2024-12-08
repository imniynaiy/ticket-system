package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
}

func (User) TableName() string {
	return "t_users"
}

type LoginResp struct {
	Token string
}
