package handlers

import (
	"net/http"

	"github.com/HsiaoCz/monster-clone/peek/services"
	"github.com/HsiaoCz/monster-clone/peek/types"
	"github.com/gofiber/fiber/v2"
)

type CommentsHandlers struct {
	cc services.CommentCaseInter
}

func NewCommentHandlers(cc services.CommentCaseInter) *CommentsHandlers {
	return &CommentsHandlers{
		cc: cc,
	}
}

func (m *CommentsHandlers) HandleCreateComment(c *fiber.Ctx) error {
	var create_comment_params types.CreateCommentParams
	if err := c.BodyParser(&create_comment_params); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	comment, err := m.cc.CreateComment(types.NewCommentFromParams(create_comment_params))
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"comment": comment,
	})
}
