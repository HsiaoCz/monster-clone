package types

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	UserID       string `json:"user_id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	Avatar       string `json:"avatar"`
	UserPassword string `json:"user_password"`
}
