package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/HsiaoCz/monster-clone/peek/db"
	"github.com/HsiaoCz/monster-clone/peek/handlers"
	"github.com/HsiaoCz/monster-clone/peek/services"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		if e, ok := err.(handlers.ErrorMsg); ok {
			return c.Status(e.Status).JSON(&e)
		}
		e := handlers.ErrorMsg{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(e.Status).JSON(&e)
	},
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	var (
		port        = os.Getenv("PORT")
		db          = db.Get()
		userCase    = services.NewUserCase(db)
		commentCase = services.NewCommentCase(db)
		postCase    = services.NewPostCase(db)

		userHandlers    = handlers.NewUserHandlers(userCase)
		postHandlers    = handlers.NewPostHandlers(postCase)
		commentHandlers = handlers.NewCommentHandlers(commentCase)
		app             = fiber.New(config)
	)

	{
		// router
		app.Post("/user", userHandlers.HandleCreateUser)

		app.Post("/post", postHandlers.HandleCreatePost)

		app.Post("/comment", commentHandlers.HandleCreateComment)
	}

	go func() {
		if err := app.Listen(port); err != nil {
			log.Fatal(err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	if err := app.Shutdown(); err != nil {
		log.Fatal(err)
	}
}
