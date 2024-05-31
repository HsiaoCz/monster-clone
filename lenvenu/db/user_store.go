package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserStorer interface {
	CreateUser(context.Context) error
}

type MongoUserStoer struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client, coll *mongo.Collection) *MongoUserStoer {
	return &MongoUserStoer{
		client: client,
		coll:   coll,
	}
}

func (u *MongoUserStoer) CreateUser(ctx context.Context) error {
	return nil
}
