package store

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type UserRedisStorer interface {
	SubscribeUser(context.Context, string, string) (string, error)
	UnSubscribeUser(context.Context, string, string) (string, error)
	BlackUser(context.Context, string, string) (string, error)
	UnBlackUser(context.Context, string, string) (string, error)
}

type UserRedisStore struct {
	db *redis.Client
}

func NewUserRedisStore(db *redis.Client) *UserRedisStore {
	return &UserRedisStore{
		db: db,
	}
}

func (u *UserRedisStore) SubscribeUser(ctx context.Context, uid string, suid string) (string, error) {
	return "", nil
}
func (u *UserRedisStore) UnSubscribeUser(ctx context.Context, uid string, suid string) (string, error) {
	return "", nil
}
func (u *UserRedisStore) BlackUser(ctx context.Context, uid string, suid string) (string, error) {
	return "", nil
}
func (u *UserRedisStore) UnBlackUser(ctx context.Context, uid string, suid string) (string, error) {
	return "", nil
}
