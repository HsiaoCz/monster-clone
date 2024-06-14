package scripts

import (
	"context"
	"os"

	"github.com/HsiaoCz/monster-clone/tony/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type postFeedStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func newPostFeedStore(ctx context.Context) (*postFeedStore, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGOURI")))
	if err != nil {
		return nil, err
	}
	return &postFeedStore{client: client, coll: client.Database(os.Getenv("DBNAME")).Collection("posts")}, nil
}

func (p *postFeedStore) createPost(ctx context.Context, post *types.Posts) (*types.Posts, error) {
	// post can repeat
	result, err := p.coll.InsertOne(ctx, post)
	if err != nil {
		return nil, err
	}
	post.ID = result.InsertedID.(primitive.ObjectID)
	return post, nil
}
