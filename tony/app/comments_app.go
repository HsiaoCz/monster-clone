package app

import (
	"net/http"

	"github.com/HsiaoCz/monster-clone/tony/store"
	"github.com/HsiaoCz/monster-clone/tony/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	createCommentsParams := types.CreateCommentsParams{}
	if err := ctx.BodyParser(&createCommentsParams); err != nil {
		return NewAPPError(http.StatusBadRequest, err.Error())
	}
	msg := createCommentsParams.Validate()
	if len(msg) != 0 {
		return ctx.Status(http.StatusBadRequest).JSON(msg)
	}
	comments := types.NewCommentFromParams(createCommentsParams)
	comment, err := c.store.CS.CreateComment(ctx.Context(), comments)
	if err != nil {
		return NewAPPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"comment": comment,
	})
}

func (c *CommentApp) HandleDeleteCommentByID(ctx *fiber.Ctx) error {
	id := ctx.Params("cid")
	cid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return NewAPPError(http.StatusBadRequest, "invalid id")
	}
	if err := c.store.CS.DeleteCommentByID(ctx.Context(), cid); err != nil {
		return NewAPPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "delete comment success",
	})
}
