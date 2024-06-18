package store

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type PostRedisStorer interface {
	CollectionPost(context.Context, string, string)
	UnCollectionPost(context.Context, string, string)
	LikePost(context.Context, string, string)
	UnLikePost(context.Context, string, string)
	WatchPosts(context.Context, string, string)
}

type PostRedisStore struct {
	db *redis.Client
}

func NewPostRedisStore(db *redis.Client) *PostRedisStore {
	return &PostRedisStore{
		db: db,
	}
}

func (p *PostRedisStore) CollectionPost(ctx context.Context, uid string, postID string)   {}
func (p *PostRedisStore) UnCollectionPost(ctx context.Context, uid string, postID string) {}
func (p *PostRedisStore) LikePost(ctx context.Context, uid string, postID string)         {}
func (p *PostRedisStore) UnLikePost(ctx context.Context, uid string, postID string)       {}
func (p *PostRedisStore) WatchPosts(ctx context.Context, uid string, postID string)       {}
