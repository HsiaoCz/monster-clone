package app

import (
	"context"
	"testing"

	"github.com/HsiaoCz/monster-clone/luccia/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	testdbUri  = "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.2.5"
	testDBName = "simon-reservation-test"
)

type testStore struct {
	client *mongo.Client
	store  *storage.Store
}

func setup(t *testing.T) *testStore {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testdbUri))
	if err != nil {
		t.Fatal(err)
	}
	userStore := storage.UserStoreInit(client, client.Database(testDBName).Collection("users"))

	return &testStore{
		client: client,
		store:  &storage.Store{Us: userStore},
	}
}

func (ts *testStore) tearDown(t *testing.T) {
	if err := ts.client.Database(testDBName).Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}
