package app

import (
	"github.com/HsiaoCz/monster-clone/tony/store"
	"github.com/gofiber/fiber/v2"
)

type SearchApp struct {
	store *store.Store
}

func NewSearchApp(store *store.Store) *SearchApp {
	return &SearchApp{
		store: store,
	}
}

func (s *SearchApp) HandleSearchUserByStr(c *fiber.Ctx) error {
	return nil
}

func (s *SearchApp) HandleSearchPostByStr(c *fiber.Ctx) error {
	return nil
}
