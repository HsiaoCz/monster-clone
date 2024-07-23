package handlers

import (
	"net/http"

	"github.com/HsiaoCz/monster-clone/peek/services"
	"github.com/HsiaoCz/monster-clone/peek/types"
	"github.com/gofiber/fiber/v2"
)

type PostHandlers struct {
	pc services.PostCaseInter
}

func NewPostHandlers(pc services.PostCaseInter) *PostHandlers {
	return &PostHandlers{
		pc: pc,
	}
}

func (p *PostHandlers) HandleCreatePost(c *fiber.Ctx) error {
	var create_post_params types.CreatePostParams
	if err := c.BodyParser(&create_post_params); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	post, err := p.pc.CreatePost(types.NewPostFromParams(create_post_params))
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": http.StatusOK,
		"post":   post,
	})
}
