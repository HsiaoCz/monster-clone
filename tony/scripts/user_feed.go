package scripts

import (
	"context"
	"errors"
	"log/slog"
	"os"

	"github.com/HsiaoCz/monster-clone/tony/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userfeed struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func newUserfeed(ctx context.Context) (*userfeed, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGOURI")))
	if err != nil {
		return nil, err
	}
	return &userfeed{client: client, coll: client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("USERCOLL"))}, nil
}

func (u *userfeed) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	filter := bson.M{
		"email":    user.Email,
		"username": user.Username,
	}
	cursor := u.coll.FindOne(ctx, filter)
	if cursor.Err() != mongo.ErrNoDocuments {
		slog.Error("db find the record error", "error message", cursor.Err())
		return nil, errors.New("the record exists")
	}
	result, err := u.coll.InsertOne(ctx, user)
	if err != nil {
		slog.Error("db insert record error", "error message", err)
		return nil, errors.New("db insert record error")
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}
