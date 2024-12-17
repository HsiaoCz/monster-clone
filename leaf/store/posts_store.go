package store

import (
	"context"

	"github.com/HsiaoCz/monster-clone/leaf/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostStorer interface {
	CreatePost(context.Context, *models.Posts) (*models.Posts, error)
	GetPostByID(context.Context, primitive.ObjectID) (*models.Posts, error)
	GetPostByTag(context.Context, primitive.ObjectID) ([]*models.Posts, error)
	GetPostByAuther(context.Context, string) ([]*models.User, error)
	GetPostByUserID(context.Context, primitive.ObjectID) ([]*models.Posts, error)
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
func (m *MongoPostStore) GetPostByID(context.Context, primitive.ObjectID) (*models.Posts, error) {
	return nil, nil
}
func (m *MongoPostStore) GetPostByTag(context.Context, primitive.ObjectID) ([]*models.Posts, error) {
	return nil, nil
}
func (m *MongoPostStore) GetPostByAuther(context.Context, string) ([]*models.User, error) {
	return nil, nil
}
func (m *MongoPostStore) GetPostByUserID(context.Context, primitive.ObjectID) ([]*models.Posts, error) {
	return nil, nil
}
