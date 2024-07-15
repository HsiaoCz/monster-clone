package data

import (
	"github.com/HsiaoCz/monster-clone/santino/types"
	"gorm.io/gorm"
)

type UserStorer interface {
	CreateUser(*types.User) (*types.User, error)
	GetUserByID(string) (*types.User, error)
}

type UserData struct {
	db *gorm.DB
}

func NewUserData(db *gorm.DB) *UserData {
	return &UserData{
		db: db,
	}
}

func (u *UserData) CreateUser(user *types.User) (*types.User, error) {
	tx := u.db.Model(&types.User{}).Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (u *UserData) GetUserByID(user_id string) (*types.User, error) {
	var user types.User
	tx := u.db.Model(&types.User{}).Where("user_id = ?", user_id).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}
