package store

import (
	"github.com/HsiaoCz/monster-clone/leaf/models"
	"gorm.io/gorm"
)

type UserStorer interface {
	CreateUser(*models.User) (*models.User, error)
}

type MySqlUserStore struct {
	db *gorm.DB
}

func NewMySqlUserStore(db *gorm.DB) *MySqlUserStore {
	return &MySqlUserStore{
		db: db,
	}
}

func (m *MySqlUserStore) CreateUser(user *models.User) (*models.User, error) {
	return nil, nil
}
