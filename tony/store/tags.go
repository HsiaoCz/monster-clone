package store

import (
	"context"
	"errors"
	"os"

	"github.com/HsiaoCz/monster-clone/tony/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TagStorer interface {
	CreateTag(context.Context, *types.Tag) (*types.Tag, error)
	GetTags(context.Context) ([]types.Tag, error)
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

func (m *MongoTagStore) GetTags(ctx context.Context) ([]types.Tag, error) {
	var tags []types.Tag
	cur, err := m.coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var tag types.Tag
		if err := cur.Decode(&tag); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}
