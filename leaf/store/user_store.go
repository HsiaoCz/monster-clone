package store

import (
	"context"
	"errors"
	"log/slog"

	"github.com/HsiaoCz/monster-clone/leaf/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserStorer interface {
	CreateUser(context.Context, *models.User) (*models.User, error)
	GetUserByID(context.Context, primitive.ObjectID) (*models.User, error)
	DeleteUserByID(context.Context, primitive.ObjectID) error
	UpdateUserByID(context.Context, primitive.ObjectID, *models.UpdateUserParmas) (*models.User, error)
	UpdateUserPassword(context.Context, primitive.ObjectID, string) error
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

func (m *MongoUserStore) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
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

func (m *MongoUserStore) GetUserByID(ctx context.Context, uid primitive.ObjectID) (*models.User, error) {
	filter := bson.D{
		{Key: "_id", Value: uid},
	}
	user := models.User{}
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

func (m *MongoUserStore) UpdateUserByID(ctx context.Context, uid primitive.ObjectID, updateUserParam *models.UpdateUserParmas) (*models.User, error) {
	filter := bson.D{
		{Key: "_id", Value: uid},
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "username", Value: updateUserParam.Username},
			{Key: "avatar", Value: updateUserParam.Avatar},
			{Key: "content", Value: updateUserParam.Content},
			{Key: "job", Value: updateUserParam.Job},
			{Key: "company", Value: updateUserParam.Company},
			{Key: "gender", Value: updateUserParam.Gender},
			{Key: "birthday", Value: updateUserParam.Birthday},
		}},
	}

	updateOptions := options.Update().SetUpsert(true)

	res, err := m.coll.UpdateOne(ctx, filter, update, updateOptions)
	if err != nil {
		return nil, err
	}

	if res.UpsertedCount == 0 {
		return nil, errors.New("no this record")
	}

	user := &models.User{}

	if err := m.coll.FindOne(ctx, filter).Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (m *MongoUserStore) UpdateUserPassword(ctx context.Context, uid primitive.ObjectID, passwd string) error {
	filter := bson.D{
		{Key: "_id", Value: uid},
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "password", Value: passwd},
		}},
	}
	updateOptions := options.Update().SetUpsert(true)

	res, err := m.coll.UpdateOne(ctx, filter, update, updateOptions)
	if err != nil {
		return err
	}

	if res.UpsertedCount == 0 {
		return errors.New("no this record")
	}
	return nil
}
