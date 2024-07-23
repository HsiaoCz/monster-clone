package services

import (
	"github.com/HsiaoCz/monster-clone/peek/types"
	"gorm.io/gorm"
)

type PostCaseInter interface {
	CreatePost(*types.Post) (*types.Post, error)
}

type PostCase struct {
	db *gorm.DB
}

func NewPostCase(db *gorm.DB) *PostCase {
	return &PostCase{
		db: db,
	}
}

func (p *PostCase) CreatePost(post *types.Post) (*types.Post, error) {
	tx := p.db.Model(&types.Post{}).Create(post)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return post, nil
}
