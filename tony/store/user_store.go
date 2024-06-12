package store

import (
	"context"

	"github.com/HsiaoCz/monster-clone/tony/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStorer interface {
	CreateUser(context.Context, *types.User) (*types.User, error)
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client, coll *mongo.Collection) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll:   coll,
	}
}

func (m *MongoUserStore) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	return nil, nil
}
