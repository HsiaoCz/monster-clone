package scripts

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

type testUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func newTestUserStore(ctx context.Context) (*testUserStore, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil {
		return nil, err
	}
	return &testUserStore{
		client: client,
		coll:   client.Database(dbname).Collection(userColl),
	}, nil
}

func (t *testUserStore) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	filter := bson.M{
		"email":    user.Email,
		"username": user.Username,
	}
	cursor := t.coll.FindOne(ctx, filter)
	if cursor.Err() != mongo.ErrNoDocuments {
		slog.Error("db find the record error", "error message", cursor.Err())
		return nil, errors.New("the record exists")
	}
	result, err := t.coll.InsertOne(ctx, user)
	if err != nil {
		slog.Error("db insert record error", "error message", err)
		return nil, errors.New("db insert record error")
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}
