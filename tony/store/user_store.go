package store

import (
	"context"
	"errors"
	"log/slog"

	"github.com/HsiaoCz/monster-clone/tony/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStorer interface {
	CreateUser(context.Context, *types.User) (*types.User, error)
	GetUserByID(context.Context, primitive.ObjectID) (*types.User, error)
	DeleteUserByID(context.Context, primitive.ObjectID) error
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client, coll *mongo.Collection) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll:   coll,
	}
}

func (m *MongoUserStore) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	filter := bson.M{
		"email":    user.Email,
		"username": user.Username,
	}
	cursor := m.coll.FindOne(ctx, filter)
	if cursor.Err() != mongo.ErrNoDocuments {
		slog.Error("db find the record error", "error message", cursor.Err())
		return nil, errors.New("the record exists")
	}
	result, err := m.coll.InsertOne(ctx, user)
	if err != nil {
		slog.Error("db insert record error", "error message", err)
		return nil, errors.New("db insert record error")
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (m *MongoUserStore) GetUserByID(ctx context.Context, uid primitive.ObjectID) (*types.User, error) {
	filter := bson.D{
		{Key: "_id", Value: uid},
	}
	user := types.User{}
	if err := m.coll.FindOne(ctx, filter).Decode(&user); err != nil {
		slog.Error("db find the record by id error", "error message", err)
		return nil, errors.New("db find the record by id error")
	}
	return &user, nil
}

func (m *MongoUserStore) DeleteUserByID(ctx context.Context, uid primitive.ObjectID) error {
	filter := bson.D{
		{Key: "_id", Value: uid},
	}
	_, err := m.coll.DeleteOne(ctx, filter)
	return err
}
