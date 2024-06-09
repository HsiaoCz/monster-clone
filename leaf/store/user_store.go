package store

import (
	"context"

	"github.com/HsiaoCz/monster-clone/leaf/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStorer interface {
	CreateUser(context.Context, *models.User) (*models.User, error)
	GetUserByID(context.Context, primitive.ObjectID) (*models.User, error)
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

func (m *MongoUserStore) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	return nil, nil
}

func (m *MongoUserStore) GetUserByID(ctx context.Context, uid primitive.ObjectID) (*models.User, error) {
	return nil, nil
}
