package app

import (
	"net/http"

	"github.com/HsiaoCz/monster-clone/tony/store"
	"github.com/gofiber/fiber/v2"
)

type UserApp struct {
	store *store.Store
}

func NewUserApp(store *store.Store) *UserApp {
	return &UserApp{
		store: store,
	}
}

func (u *UserApp) HandleCreateUser(c *fiber.Ctx) error {
	// return c.Status(http.StatusOK).JSON(fiber.Map{
	// 	"message": "something",
	// })
	return NewAPPError(http.StatusInternalServerError, "Hello my man")
}
