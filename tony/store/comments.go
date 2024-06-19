package store

import (
	"context"
	"os"

	"github.com/HsiaoCz/monster-clone/tony/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentStorer interface {
	CreateComment(context.Context, *types.Comments) (*types.Comments, error)
	DeleteCommentByID(context.Context, primitive.ObjectID) error
	GetCommentsByPostID(context.Context, primitive.ObjectID) ([]*types.Comments, error)
	GetCommentsByParentID(context.Context, primitive.ObjectID) ([]*types.Comments, error)
}

type MongoCommentStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoCommentStore(client *mongo.Client, coll *mongo.Collection) *MongoCommentStore {
	return &MongoCommentStore{
		client: client,
		coll:   client.Database(os.Getenv("DBNAME")).Collection("comments"),
	}
}

func (m *MongoCommentStore) CreateComment(ctx context.Context, comment *types.Comments) (*types.Comments, error) {
	return nil, nil
}
func (m *MongoCommentStore) DeleteCommentByID(ctx context.Context, cid primitive.ObjectID) error {
	return nil
}

func (m *MongoCommentStore) GetCommentsByPostID(ctx context.Context, pid primitive.ObjectID) ([]*types.Comments, error) {
	return nil, nil
}

func (m *MongoCommentStore) GetCommentsByParentID(ctx context.Context, parentID primitive.ObjectID) ([]*types.Comments, error) {
	return nil, nil
}
