package store

import (
	"context"

	"github.com/HsiaoCz/monster-clone/leaf/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostStorer interface {
	CreatePost(context.Context, *models.Posts) (*models.Posts, error)
	GetPostByID(context.Context, primitive.ObjectID) (*models.Posts, error)
	GetPostByTag(context.Context, string) ([]models.Posts, error)
	GetPostByAuther(context.Context, string) ([]models.Posts, error)
	GetPostByUserID(context.Context, primitive.ObjectID) ([]models.Posts, error)
}

type MongoPostStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoPostStore(client *mongo.Client, coll *mongo.Collection) *MongoPostStore {
	return &MongoPostStore{
		client: client,
		coll:   coll,
	}
}

func (m *MongoPostStore) CreatePost(ctx context.Context, post *models.Posts) (*models.Posts, error) {
	// post can muti create
	// so don't need to check
	res, err := m.coll.InsertOne(ctx, post)
	if err != nil {
		return nil, err
	}
	post.ID = res.InsertedID.(primitive.ObjectID)
	return post, nil
}
func (m *MongoPostStore) GetPostByID(ctx context.Context, post_id primitive.ObjectID) (*models.Posts, error) {
	filter := bson.D{
		{Key: "_id", Value: post_id},
	}
	var post models.Posts
	if err := m.coll.FindOne(ctx, filter).Decode(&post); err != nil {
		return nil, err
	}
	return &post, nil
}
func (m *MongoPostStore) GetPostByTag(ctx context.Context, tag string) ([]models.Posts, error) {
	filter := bson.D{
		{Key: "tags", Value: bson.D{{Key: "$in", Value: []string{tag}}}},
	}
	var posts []models.Posts
	cur, err := m.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	if err := cur.All(ctx, &posts); err != nil {
		return nil, err
	}
	return posts, nil
}
func (m *MongoPostStore) GetPostByAuther(context.Context, string) ([]models.Posts, error) {
	return nil, nil
}
func (m *MongoPostStore) GetPostByUserID(context.Context, primitive.ObjectID) ([]models.Posts, error) {
	return nil, nil
}
