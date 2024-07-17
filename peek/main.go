package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/HsiaoCz/monster-clone/peek/db"
	"github.com/HsiaoCz/monster-clone/peek/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
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

	var (
		port = os.Getenv("PORT")
		app  = fiber.New(config)
	)
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
