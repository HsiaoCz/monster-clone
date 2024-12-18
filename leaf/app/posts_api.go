package app

import (
	"net/http"

	"github.com/HsiaoCz/monster-clone/leaf/models"
	"github.com/HsiaoCz/monster-clone/leaf/store"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	createPostParam := models.CreatePostsParams{}
	if err := c.BodyParser(&createPostParam); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}

	// p.store.Post.CreatePost(c.Context(),)
	return nil
}
func (p *PostApp) HandleDeletePost(c *fiber.Ctx) error {
	id := c.Params("pid")
	post_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	p.store.Post.GetPostByID(c.Context(), post_id)
	return nil
}
func (p *PostApp) HandleGetPostsByID(c *fiber.Ctx) error {
	id := c.Params("pid")
	post_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	p.store.Post.GetPostByID(c.Context(), post_id)
	return nil
}
func (p *PostApp) HandleCreatePostByClassify(c *fiber.Ctx) error {
	return nil
}
