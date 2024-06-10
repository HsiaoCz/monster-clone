package store

import (
	"context"
	"errors"

	"github.com/HsiaoCz/monster-clone/leaf/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TagStorer interface {
	CreateTags(context.Context, *models.Tag) (*models.Tag, error)
}

type MongoTagStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoTagStore(client *mongo.Client, coll *mongo.Collection) *MongoTagStore {
	return &MongoTagStore{
		client: client,
		coll:   coll,
	}
}

func (m *MongoTagStore) CreateTags(ctx context.Context, tag *models.Tag) (*models.Tag, error) {
	filter := bson.D{
		{Key: "content", Value: tag.Content},
	}
	cursor := m.coll.FindOne(ctx, filter)
	if cursor.Err() != mongo.ErrNoDocuments {
		return nil, errors.New("this record exists")
	}
	result, err := m.coll.InsertOne(ctx, tag)
	if err != nil {
		return nil, err
	}
	tag.ID = result.InsertedID.(primitive.ObjectID)
	return tag, nil
}
