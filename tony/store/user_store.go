package store

import (
	"context"
	"errors"
	"log/slog"

	"github.com/HsiaoCz/monster-clone/tony/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserStorer interface {
	CreateUser(context.Context, *types.User) (*types.User, error)
	GetUserByID(context.Context, primitive.ObjectID) (*types.User, error)
	DeleteUserByID(context.Context, primitive.ObjectID) error
	UpdateUserByID(context.Context, primitive.ObjectID, *types.UpdateUserParmas) (*types.User, error)
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

func (m *MongoUserStore) UpdateUserByID(ctx context.Context, uid primitive.ObjectID, updateUserParma *types.UpdateUserParmas) (*types.User, error) {
	filter := bson.D{
		{Key: "_id", Value: uid},
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "", Value: ""},
			{Key: "", Value: ""},
			{Key: "", Value: ""},
			{Key: "", Value: ""},
			{Key: "", Value: ""},
		}},
	}

	updateOptions := options.Update().SetUpsert(true)

	_, err := m.coll.UpdateOne(ctx, filter, update, updateOptions)
	if err != nil {
		return nil, err
	}

	res := &types.User{}

	if err := m.coll.FindOne(ctx, filter).Decode(res); err != nil {
		return nil, err
	}

	return res, nil
}
