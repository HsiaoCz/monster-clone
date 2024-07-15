package data

import (
	"github.com/HsiaoCz/monster-clone/santino/types"
	"gorm.io/gorm"
)

type UserStorer interface {
	CreateUser(*types.User) (*types.User, error)
	GetUserByID(string) (*types.User, error)
	UpdateUserByID(string, *types.UpdateUser) (*types.User, error)
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

func (u *UserData) UpdateUserByID(user_id string, update_user *types.UpdateUser) (*types.User, error) {
	var user types.User
	tx := u.db.Model(&types.User{}).Where("user_id = ?", user_id).Updates(map[string]any{"username": update_user.Username, "avatar": update_user.Avatar, "synopsis": update_user.Synopsis, "background_image": update_user.Background_Image})
	if tx.Error != nil {
		return nil, tx.Error
	}
	tx1 := u.db.Model(&types.User{}).Find(&user, "user_id = ?", user_id)
	if tx1.Error != nil {
		return nil, tx1.Error
	}
	return &user, nil
}
