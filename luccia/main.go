package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/HsiaoCz/monster-clone/luccia/app"
	"github.com/HsiaoCz/monster-clone/luccia/store"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGOURL")))
	if err != nil {
		log.Fatal(err)
	}

	go func(ctx context.Context) {
		if err := client.Ping(ctx, &readpref.ReadPref{}); err != nil {
			log.Fatal(err)
		}
	}(ctx)

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	var (
		userColl  = client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("USERCOLL"))
		userStore = store.UserStoreInit(client, userColl)
		userApp   = app.UserAppInit(userStore)
		router    = http.NewServeMux()
	)

	{
		// router
		router.HandleFunc("POST /user", app.TransferHandlerfunc(userApp.HandleCreateUser))
	}

	logrus.WithFields(logrus.Fields{
		"listen address": os.Getenv("PORT"),
	}).Info("the http server is running")

	http.ListenAndServe(os.Getenv("PORT"), router)
}
