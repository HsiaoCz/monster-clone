package handler

import (
	"github.com/HsiaoCz/monster-clone/santino/data"
	"github.com/gofiber/fiber/v2"
)

type PostHandler struct {
	post data.PostStorer
}

func NewPostHandler(post data.PostStorer) *PostHandler {
	return &PostHandler{
		post: post,
	}
}

func (p *PostHandler) HandleCreatePost(c *fiber.Ctx) error {
	return nil
}
