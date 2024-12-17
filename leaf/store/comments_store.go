package store

import (
	"context"

	"github.com/HsiaoCz/monster-clone/leaf/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentStorer interface {
	CreateComments(context.Context, *models.Comments) (*models.Comments, error)
	UpdateComment(context.Context, int) (*models.Comments, error)
}

type MongoCommentStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoCommentStore(client *mongo.Client, coll *mongo.Collection) *MongoCommentStore {
	return &MongoCommentStore{
		client: client,
		coll:   coll,
	}
}

func (m *MongoCommentStore) CreateComments(ctx context.Context, comment *models.Comments) (*models.Comments, error) {
	return nil, nil
}

func (m *MongoCommentStore) UpdateComment(ctx context.Context, likes int) (*models.Comments, error) {
	return nil, nil
}
