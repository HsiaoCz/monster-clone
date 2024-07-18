package data

import (
	"github.com/HsiaoCz/monster-clone/santino/types"
	"gorm.io/gorm"
)

type AdminStorer interface {
	CreateAdmin(*types.Admin) (*types.Admin, error)
	DeleteAdminByID(string) error
}

type AdminStore struct {
	db *gorm.DB
}

func NewAdminStore(db *gorm.DB) *AdminStore {
	return &AdminStore{
		db: db,
	}
}

func (a *AdminStore) CreateAdmin(admin *types.Admin) (*types.Admin, error) {
	return admin, nil
}

func (a *AdminStore) DeleteAdminByID(user_id string) error {
	return a.db.Where("user_id = ?", user_id).Delete(&types.Admin{}).Error
}
