package store

import (
	"github.com/HsiaoCz/monster-clone/leaf/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStorer interface {
	CreateUser(*models.User) (*models.User, error)
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

func (m *MongoUserStore) CreateUser(user *models.User) (*models.User, error) {
	return nil, nil
}
