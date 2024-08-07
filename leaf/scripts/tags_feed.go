package scripts

import (
	"context"
	"errors"

	"github.com/HsiaoCz/monster-clone/leaf/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type testTagStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func newTestTagStore(ctx context.Context) (*testTagStore, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil {
		return nil, err
	}
	return &testTagStore{
		client: client,
		coll:   client.Database(dbname).Collection(tagColl),
	}, nil
}

func (t *testTagStore) CreateTags(ctx context.Context, tag *models.Tag) (*models.Tag, error) {
	filter := bson.D{
		{Key: "content", Value: tag.Content},
	}
	cursor := t.coll.FindOne(ctx, filter)
	if cursor.Err() != mongo.ErrNoDocuments {
		return nil, errors.New("this record exists")
	}
	result, err := t.coll.InsertOne(ctx, tag)
	if err != nil {
		return nil, err
	}
	tag.ID = result.InsertedID.(primitive.ObjectID)
	return tag, nil
}
