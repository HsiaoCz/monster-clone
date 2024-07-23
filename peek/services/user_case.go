package services

import (
	"github.com/HsiaoCz/monster-clone/peek/types"
	"gorm.io/gorm"
)

type UserCaseInter interface {
	CreateUser(*types.User) (*types.User, error)
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
