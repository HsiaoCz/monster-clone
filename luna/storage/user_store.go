package storage

import (
	"context"
	"errors"

	"github.com/HsiaoCz/monster-clone/luna/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (u *UserStore) GetUserByEmail(ctx context.Context, email string) (*types.Users, error) {
	var user types.Users
	filter := bson.D{
		{Key: "email", Value: email},
	}
	if err := u.coll.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserStore) UpdateUser(ctx context.Context, uid primitive.ObjectID, updateParams *types.UpdateUserParams) (*types.Users, error) {
	filter := bson.D{{Key: "_id", Value: uid}}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "username", Value: updateParams.Username},
			{Key: "content", Value: updateParams.Content},
			{Key: "company", Value: updateParams.Company},
			{Key: "job", Value: updateParams.Job},
			{Key: "avatar", Value: updateParams.Avatar},
			{Key: "gender", Value: updateParams.Gender},
		}},
	}
	updateOptions := options.Update().SetUpsert(true)
	_, err := u.coll.UpdateOne(ctx, filter, update, updateOptions)
	if err != nil {
		return nil, err
	}
	var user types.Users
	if err := u.coll.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
