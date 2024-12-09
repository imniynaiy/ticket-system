package model

type UserSession struct {
	UserID uint `json:"user_id"`
	Role   int  `json:"role"`
}
