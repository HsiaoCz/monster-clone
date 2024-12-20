package store

import (
	"context"
	"os"

	"github.com/HsiaoCz/monster-clone/tony/types"
	"go.mongodb.org/mongo-driver/bson"
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
	cur, err := m.coll.InsertOne(ctx, post)
	if err != nil {
		return nil, err
	}
	post.ID = cur.InsertedID.(primitive.ObjectID)
	return post, nil
}
func (m *MongoPostStore) DeletePost(ctx context.Context, pid primitive.ObjectID) error {
	filter := bson.D{
		{Key: "_id", Value: pid},
	}
	_, err := m.coll.DeleteOne(ctx, filter)
	return err
}
func (m *MongoPostStore) GetPosts(ctx context.Context) ([]*types.Posts, error) {
	var posts []*types.Posts
	cur, err := m.coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var post types.Posts
		if err := cur.Decode(&post); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	return posts, nil
}
