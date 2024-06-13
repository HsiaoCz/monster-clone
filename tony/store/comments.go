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
	DeleteComment(context.Context, primitive.ObjectID) error
	GetComments(context.Context) ([]*types.Comments, error)
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
func (m *MongoCommentStore) DeleteComment(ctx context.Context, id primitive.ObjectID) (*types.Comments, error) {
	return nil, nil
}

func (m *MongoCommentStore) GetComments(ctx context.Context) ([]*types.Comments, error) {
	return nil, nil
}
