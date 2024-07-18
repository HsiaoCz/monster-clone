package handler

import (
	"net/http"

	"github.com/HsiaoCz/monster-clone/santino/data"
	"github.com/HsiaoCz/monster-clone/santino/types"
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
	var create_post types.CreatePostParams
	if err := c.BodyParser(&create_post); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	post, err := p.post.CreatePost(types.NewPostFromParams(create_post))
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "create post success",
		"post":    post,
	})
}
