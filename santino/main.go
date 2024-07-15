package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/HsiaoCz/monster-clone/santino/data"
	"github.com/HsiaoCz/monster-clone/santino/db"
	"github.com/HsiaoCz/monster-clone/santino/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		if e, ok := err.(handler.ErrorMsg); ok {
			return c.Status(e.Status).JSON(&e)
		}
		e := handler.ErrorMsg{
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
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	var (
		port           = os.Getenv("PORT")
		userData       = data.NewUserData(db.Get())
		userHandler    = handler.NewUserHandler(userData)
		postHandler    = &handler.PostHandler{}
		adminHandler   = &handler.AdminHandler{}
		commentHandler = &handler.CommentHandler{}
		app            = fiber.New(config)
	)
	{
		// router
		app.Post("/user", userHandler.HandleCreateUser)
		app.Get("/user/:user_id", userHandler.HandleGetUserByID)
		app.Post("/post", postHandler.HandleCreatePost)
		app.Post("/admin", adminHandler.HandleCreateAdmin)
		app.Post("/comment", commentHandler.HandleCreateComment)
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
