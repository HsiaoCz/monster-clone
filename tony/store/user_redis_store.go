package store

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type UserRedisStorer interface {
	SubscribeUser(context.Context, string, string)
	UnSubscribeUser(context.Context, string, string)
	BlackUser(context.Context, string, string)
}

type UserRedisStore struct {
	db *redis.Client
}

func NewUserRedisStore(db *redis.Client) *UserRedisStore {
	return &UserRedisStore{
		db: db,
	}
}

func (u *UserRedisStore) SubscribeUser(ctx context.Context, uid string, suid string)   {}
func (u *UserRedisStore) UnSubscribeUser(ctx context.Context, uid string, suid string) {}
func (u *UserRedisStore) BlackUser(ctx context.Context, uid string, suid string)       {}
