package store

import (
	"context"
	"os"

	"github.com/HsiaoCz/monster-clone/tony/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type TagStorer interface {
	CreateTag(context.Context, *types.Tag) (*types.Tag, error)
	GetTags(context.Context) ([]*types.Tag, error)
}

type MongoTagStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoTagStore(client *mongo.Client, coll *mongo.Collection) *MongoTagStore {
	return &MongoTagStore{
		client: client,
		coll:   client.Database(os.Getenv("DBNAME")).Collection("tags"),
	}
}

func (m *MongoTagStore) CreateTag(ctx context.Context, tag *types.Tag) (*types.Tag, error) {
	return nil, nil
}

func (m *MongoTagStore) GetTags(ctx context.Context, tag *types.Tag) (*types.Tag, error) {
	return nil, nil
}
