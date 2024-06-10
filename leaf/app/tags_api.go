package app

import (
	"github.com/HsiaoCz/monster-clone/leaf/store"
	"github.com/gofiber/fiber/v2"
)

type TagsApp struct {
	store *store.Store
}

func NewTagsApp(store *store.Store) *TagsApp {
	return &TagsApp{
		store: store,
	}
}

func (t *TagsApp) HandleCreateTags(c *fiber.Ctx) error {
	return nil
}

func (t *TagsApp) HandleDeleteTags(c *fiber.Ctx) error {
	return nil
}

func (t *TagsApp) HandleGetTags(c *fiber.Ctx) error {
	return nil
}
