package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/HsiaoCz/monster-clone/leaf/app"
	"github.com/HsiaoCz/monster-clone/leaf/conf"
	"github.com/HsiaoCz/monster-clone/leaf/store"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		if e, ok := err.(app.APIError); ok {
			return c.Status(e.Status).JSON(&e)
		}
		aErr := app.APIError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(aErr.Status).JSON(&aErr)
	},
}

func main() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal(err)
	// }
	// env is totally shit
	// var (
	// 	user     = os.Getenv("MUSER")
	// 	password = os.Getenv("PASSWORD")
	// 	host     = os.Getenv("HOST")
	// 	port     = os.Getenv("PORT")
	// 	dbname   = os.Getenv("DBNAME")
	// )
	// we don't need this shit
	// but if the db connect error we use this to check shit where?
	// fmt.Println(user, password, host, port, dbname)
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	if err := conf.ParseConfig(); err != nil {
		slog.Error("parse config error", "error message", err)
		return
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(conf.Conf.App.MongoURI))
	if err != nil {
		slog.Error("mongo db connect error", "error message", err)
		return
	}

	var (
		logger            = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
		userColl          = client.Database(conf.Conf.App.DBname).Collection(conf.Conf.App.UserColl)
		postColl          = client.Database(conf.Conf.App.DBname).Collection(conf.Conf.App.PostColl)
		commentColl       = client.Database(conf.Conf.App.DBname).Collection(conf.Conf.App.CommentColl)
		tagColl           = client.Database(conf.Conf.App.DBname).Collection(conf.Conf.App.TagsColl)
		mongoUserStore    = store.NewMongoUserStore(client, userColl)
		mongoPostStore    = store.NewMongoPostStore(client, postColl)
		mongoCommentStore = store.NewMongoCommentStore(client, commentColl)
		mongoTagStore     = store.NewMongoTagStore(client, tagColl)
		store             = &store.Store{User: mongoUserStore, Tag: mongoTagStore, Comment: mongoCommentStore, Post: mongoPostStore}
		userHandlers      = app.NewUserAPI(store)
		tagHandlers       = app.NewTagsApp(store)
		commentHandlers   = app.NewCommentsApp(store)
		postHandlers      = app.NewPostApp(store)
		router            = fiber.New(config)
		av1               = router.Group("/app/v1")
	)
	slog.SetDefault(logger)
	// routers
	{
		av1.Post("/user", userHandlers.HandleCreateUser)
		av1.Get("/user/:uid", userHandlers.HandleGetUserByID)
		av1.Delete("/user/:id", userHandlers.HandleDeleteUserByID)
		av1.Post("/user/update", userHandlers.HandleUpdateUser)

		av1.Post("/post", postHandlers.HandleCreatePost)
		av1.Post("/tag", tagHandlers.HandleCreateTags)
		av1.Post("/comment", commentHandlers.HandleCreateComments)
	}

	// restart and shutdown
	go func() {
		if err := router.Listen(conf.Conf.App.Port); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	if err := router.Shutdown(); err != nil {
		panic(err)
	}
}
