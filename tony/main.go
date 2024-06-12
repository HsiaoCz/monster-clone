package main

import (
	"log"
	"net/http"

	"github.com/HsiaoCz/monster-clone/tony/app"
	"github.com/HsiaoCz/monster-clone/tony/config"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
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
	router := fiber.New(errConfig)
	router.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "ok",
		})
	})
	router.Listen(config.GetPort())
}
