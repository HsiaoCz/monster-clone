package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/HsiaoCz/monster-clone/tony/app"
	"github.com/HsiaoCz/monster-clone/tony/store"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var errConfig = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		if e, ok := err.(app.APPError); ok {
			return c.Status(e.Status).JSON(&e)
		}
		appError := app.APPError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(appError.Status).JSON(&appError)
	},
}

func main() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGOURI")))
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := client.Ping(ctx, &readpref.ReadPref{}); err != nil {
			log.Fatal(err)
		}
	}()

	var (
		userColl       = client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("USERCOLL"))
		mongoUserStore = store.NewMongoUserStore(client, userColl)
		store          = &store.Store{US: mongoUserStore}
		userHandler    = app.NewUserApp(store)
		router         = fiber.New(errConfig)
		v1             = router.Group("/app/v1")
	)

	// router
	{
		v1.Post("/user", userHandler.HandleCreateUser)
		v1.Get("/user/:uid", userHandler.HandleGetUserByID)
		v1.Delete("/user/:uid", userHandler.HandleDeleteUserByID)
		v1.Post("/user/:uid", userHandler.HandleUpdateUserByID)
	}

	// restart and shutdown
	go func() {
		if err := router.Listen(os.Getenv("PORT")); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	if err := router.Shutdown(); err != nil {
		log.Fatal(err)
	}
}
