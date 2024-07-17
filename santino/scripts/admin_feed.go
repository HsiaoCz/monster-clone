package scripts

import (
	"github.com/HsiaoCz/monster-clone/santino/types"
	"gorm.io/gorm"
)

type AdminFeed struct {
	db *gorm.DB
}

func NewAdminFeed(db *gorm.DB) *AdminFeed {
	return &AdminFeed{
		db: db,
	}
}

func (a *AdminFeed) CreateAdmin(admin *types.Admin) (*types.Admin, error) {
	tx := a.db.Model(&types.Admin{}).Create(admin)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return admin, nil
}
