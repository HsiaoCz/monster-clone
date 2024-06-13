package app

import (
	"github.com/HsiaoCz/monster-clone/tony/store"
	"github.com/gofiber/fiber/v2"
)

type PostApp struct {
	store *store.Store
}

func NewPostApp(store *store.Store) *PostApp {
	return &PostApp{
		store: store,
	}
}

func (p *PostApp) HandleCreatePost(c *fiber.Ctx) error {
	return nil
}
