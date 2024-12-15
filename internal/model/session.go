package model

type UserSession struct {
	UserID   uint `json:"user_id"`
	IsAdmin  bool `json:"is_admin"`
}
