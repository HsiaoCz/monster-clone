package storage

import (
	"context"
	"errors"

	"github.com/HsiaoCz/monster-clone/luna/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStoreInter interface {
	CreateUser(context.Context, *types.Users) (*types.Users, error)
}

type UserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func UserStoreInit(client *mongo.Client, coll *mongo.Collection) *UserStore {
	return &UserStore{
		client: client,
		coll:   coll,
	}
}

func (u *UserStore) CreateUser(ctx context.Context, user *types.Users) (*types.Users, error) {
	filter := bson.M{
		"email":    user.Email,
		"username": user.Username,
	}
	cursor := u.coll.FindOne(ctx, filter)
	if cursor.Err() != mongo.ErrNoDocuments {
		return nil, errors.New("the record exists")
	}
	result, err := u.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, errors.New("db insert record error")
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (u *UserStore) GetUserByID(ctx context.Context, uid primitive.ObjectID) (*types.Users, error) {
	var user types.Users
	filter := bson.D{
		{Key: "_id", Value: uid},
	}
	if err := u.coll.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, errors.New("db find the record by id error")
	}
	return &user, nil
}

func (u *UserStore) DeleteUserByID(ctx context.Context, uid primitive.ObjectID) error {
	filter := bson.D{
		{Key: "_id", Value: uid},
	}
	_, err := u.coll.DeleteOne(ctx, filter)
	return err
}
