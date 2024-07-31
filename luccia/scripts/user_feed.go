package scripts

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserFeed struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func UserFeedInit() (*UserFeed, error) {
	if err := godotenv.Load("../.env"); err != nil {
		return nil, err
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGOURL")))
	if err != nil {
		return nil, err
	}
	return &UserFeed{
		client: client,
		coll:   client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("USERCOLL")),
	}, nil
}
