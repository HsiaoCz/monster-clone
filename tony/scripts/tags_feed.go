package scripts

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/HsiaoCz/monster-clone/tony/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type tagTestStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func newTagTestStore(ctx context.Context) (*tagTestStore, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGOURI")))
	if err != nil {
		log.Fatal(err)
	}
	return &tagTestStore{
		client: client,
		coll:   client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("TAGSCOLL")),
	}, nil
}

func (t *tagTestStore) CreateTags(ctx context.Context, tag *types.Tag) (*types.Tag, error) {
	filter := bson.D{
		{Key: "content", Value: tag.Content},
	}
	res := t.coll.FindOne(ctx, filter)
	if res.Err() != mongo.ErrNoDocuments {
		return nil, errors.New("the record exists")
	}
	result, err := t.coll.InsertOne(ctx, tag)
	if err != nil {
		return nil, err
	}
	tag.ID = result.InsertedID.(primitive.ObjectID)
	return tag, nil
}

func (t *tagTestStore) GetTags(ctx context.Context) ([]*types.Tag, error) {
	tags := []*types.Tag{}
	cur, err := t.coll.Find(ctx, bson.M{})
	if err != nil {
		return []*types.Tag{}, err
	}
	if err := cur.All(ctx, &tags); err != nil {
		return []*types.Tag{}, err
	}
	return tags, nil
}
