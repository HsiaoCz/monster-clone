package store

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type CommentRedisStorer interface {
	LikeComment(context.Context, string)
	UnlikeComment(context.Context, string)
}

type CommentRedisStore struct {
	db *redis.Client
}

func NewCommentRedisStore(db *redis.Client) *CommentRedisStore {
	return &CommentRedisStore{
		db: db,
	}
}

func (c *CommentRedisStore) LikeComment(ctx context.Context, commentID string)   {}
func (c *CommentRedisStore) UnlikeComment(ctx context.Context, commentID string) {}
