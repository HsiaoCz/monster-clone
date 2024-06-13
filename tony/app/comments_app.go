package app

import (
	"github.com/HsiaoCz/monster-clone/leaf/store"
	"github.com/gofiber/fiber/v2"
)

type CommentApp struct {
	store *store.Store
}

func NewCommentApp(store *store.Store) *CommentApp {
	return &CommentApp{
		store: store,
	}
}

func (c *CommentApp) HandleCreateComment(ctx *fiber.Ctx) error {
	return nil
}
