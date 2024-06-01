package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID   string `gorm:"column:user_id;type:varchar(100);" json:"user_id"`
	Username string `gorm:"column:username;type:varchar(40);" json:"username"`
	Password string `gorm:"column:password;type:varchar(100);" json:"-"`
	Email    string `gorm:"column:email;type:varchar(60);" json:"email"`
}
