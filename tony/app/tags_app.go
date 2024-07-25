package app

import (
	"github.com/HsiaoCz/monster-clone/tony/store"
	"github.com/gofiber/fiber/v2"
)

type TagApp struct {
	store *store.Store
}

func NewTagApp(store *store.Store) *TagApp {
	return &TagApp{
		store: store,
	}
}

func (t *TagApp) HandleCreateTag(c *fiber.Ctx) error {
	return nil
}

func (t *TagApp)HandleDeleteTag(c *fiber.Ctx)error{
	return nil
}