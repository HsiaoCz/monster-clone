package store

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	test_mongoUrl = ""
	test_dbname   = "luccia_test"
)

var client *mongo.Client

func Init() error {
	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI(test_mongoUrl))
	if err != nil {
		return err
	}
	client = c
	return nil
}
