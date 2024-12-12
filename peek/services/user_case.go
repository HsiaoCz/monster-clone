package services

import (
	"github.com/HsiaoCz/monster-clone/peek/types"
	"gorm.io/gorm"
)

type UserCaseInter interface {
	CreateUser(*types.User) (*types.User, error)
	GetUserByID(string) (*types.User, error)
	DeleteUserByID(string) error
	GetUserByEmailAndPassword(types.UserLoginParmas) (*types.User, error)
}

type UserCase struct {
	db *gorm.DB
}

func NewUserCase(db *gorm.DB) *UserCase {
	return &UserCase{
		db: db,
	}
}

func (u *UserCase) CreateUser(user *types.User) (*types.User, error) {
	tx := u.db.Model(&types.User{}).Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (u *UserCase) GetUserByID(user_id string) (*types.User, error) {
	var user types.User
	tx := u.db.Debug().Model(&types.User{}).Where("user_id = ?", user_id).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}

func (u *UserCase) DeleteUserByID(user_id string) error {
	var user types.User
	tx := u.db.Debug().Model(types.User{}).Where("user_id = ?", user_id).Delete(&user)
	return tx.Error
}

func (u *UserCase) GetUserByEmailAndPassword(parmas types.UserLoginParmas) (*types.User, error) {
	var user types.User
	tx := u.db.Debug().Model(&types.User{}).Where("email = ? AND password", parmas.Email, parmas.Password).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &user, nil
}
