package app

import (
	"net/http"

	"github.com/HsiaoCz/monster-clone/leaf/models"
	"github.com/HsiaoCz/monster-clone/leaf/store"
	"github.com/gofiber/fiber/v2"
)

type CommentsApp struct {
	store *store.Store
}

func NewCommentsApp(store *store.Store) *CommentsApp {
	return &CommentsApp{
		store: store,
	}
}

func (m *CommentsApp) HandleCreateComments(c *fiber.Ctx) error {
	var comment models.CreateCommentsParams
	if err := c.BodyParser(&comment); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	return nil
}
func (m *CommentsApp) HandleGetCommentsByID(c *fiber.Ctx) error {
	return nil
}
func (m *CommentsApp) HandleDeleteComments(c *fiber.Ctx) error {
	return nil
}
