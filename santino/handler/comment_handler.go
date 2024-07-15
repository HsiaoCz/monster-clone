package handler

import (
	"github.com/HsiaoCz/monster-clone/santino/data"
	"github.com/gofiber/fiber/v2"
)

type CommentHandler struct {
	comment data.CommentStorer
}

func NewCommentHandler(comment data.CommentStorer) *CommentHandler {
	return &CommentHandler{
		comment: comment,
	}
}

func (m *CommentHandler) HandleCreateComment(c *fiber.Ctx) error {
	return nil
}
