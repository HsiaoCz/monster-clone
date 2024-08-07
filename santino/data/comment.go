package data

import (
	"github.com/HsiaoCz/monster-clone/santino/types"
	"gorm.io/gorm"
)

type CommentStorer interface {
	CreateComment(*types.Comment) (*types.Comment, error)
	DeleteCommentByID(string) error
}

type CommentStore struct {
	db *gorm.DB
}

func NewCommentStore(db *gorm.DB) *CommentStore {
	return &CommentStore{
		db: db,
	}
}

func (c *CommentStore) CreateComment(comment *types.Comment) (*types.Comment, error) {
	tx := c.db.Model(&types.Comment{}).Create(comment)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return comment, nil
}

func (c *CommentStore) DeleteCommentByID(comment_id string) error {
	return c.db.Where("comment_id = ?", comment_id).Delete(&types.Comment{}).Error
}
