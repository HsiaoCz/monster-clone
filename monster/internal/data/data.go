package data

import (
	"context"

	"github.com/HsiaoCz/monster-clone/monster/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	client *mongo.Client
	coll   *mongo.Collection
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(c.Database.MongoUrl))
	if err != nil {
		logger.Log(log.LevelFatal, "mongo database connect error", err)
	}
	go func() {
		if err := client.Ping(ctx, &readpref.ReadPref{}); err != nil {
			logger.Log(log.LevelFatal, "mongo database ping error", err)
		}
	}()
	return &Data{
		client: client,
		coll:   client.Database(c.Database.Dbname).Collection(c.Database.Coll),
	}, cleanup, nil
}
