package services

import (
	"github.com/HsiaoCz/monster-clone/peek/types"
	"gorm.io/gorm"
)

type CommentCaseInter interface {
	CreateComment(*types.Comment) (*types.Comment, error)
}

type CommentCase struct {
	db *gorm.DB
}

func NewCommentCase(db *gorm.DB) *CommentCase {
	return &CommentCase{
		db: db,
	}
}

func (c *CommentCase) CreateComment(comment *types.Comment) (*types.Comment, error) {
	tx := c.db.Model(&types.Comment{}).Create(comment)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return comment, nil
}
