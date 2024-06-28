package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/HsiaoCz/monster-clone/monster/gateway/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		if e, ok := err.(handlers.ErrorMsg); ok {
			return c.Status(e.Status).JSON(&e)
		}
		errMsg := handlers.ErrorMsg{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(errMsg.Status).JSON(&errMsg)
	},
}

func main() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatal(err)
	}

	var (
		port         = os.Getenv("PORT")
		postHandlers = &handlers.PostHandlers{}
		router       = fiber.New(config)
		v1           = router.Group("/api/v1")
	)

	{
		v1.Post("/post", postHandlers.HandleCreatePost)
		v1.Delete("/post/:pid", postHandlers.HandleDeletePostByID)
	}

	// restart and shutdown
	go func() {
		if err := router.Listen(port); err != nil {
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
