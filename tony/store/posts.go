package store

import (
	"context"
	"os"

	"github.com/HsiaoCz/monster-clone/tony/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostStorer interface {
	CreatePost(context.Context, *types.Posts) (*types.Posts, error)
	DeletePost(context.Context, primitive.ObjectID) error
	GetPosts(context.Context) ([]*types.Posts, error)
}

type MongoPostStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoPostStore(client *mongo.Client, coll *mongo.Collection) *MongoPostStore {
	return &MongoPostStore{
		client: client,
		coll:   client.Database(os.Getenv("DBNAME")).Collection("posts"),
	}
}

func (m *MongoPostStore) CreatePost(ctx context.Context, post *types.Posts) (*types.Posts, error) {
	return nil, nil
}
func (m *MongoPostStore) DeletePost(ctx context.Context, pid primitive.ObjectID) error {
	return nil
}
func (m *MongoPostStore) GetPosts(ctx context.Context) ([]*types.Posts, error) {
	return nil, nil
}
