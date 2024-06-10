package app

import (
	"github.com/HsiaoCz/monster-clone/leaf/store"
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
func (p *PostApp) HandleDeletePost(c *fiber.Ctx) error {
	return nil
}
func (p *PostApp) HandleGetPostsByID(c *fiber.Ctx) error {
	return nil
}
func (p *PostApp) HandleCreatePostByClassify(c *fiber.Ctx) error {
	return nil
}
