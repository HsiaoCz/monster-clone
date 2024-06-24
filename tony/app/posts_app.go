package app

import (
	"net/http"

	"github.com/HsiaoCz/monster-clone/tony/app/middleware"
	"github.com/HsiaoCz/monster-clone/tony/store"
	"github.com/HsiaoCz/monster-clone/tony/types"
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
	userInfo, ok := c.UserContext().Value(middleware.CtxUserInfoKey).(types.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusNonAuthoritativeInfo, "user need login")
	}
	var createPostParam types.CreatePostsParams
	if err := c.BodyParser(&createPostParam); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	msg := createPostParam.Validate()
	if len(msg) != 0 {
		return c.Status(http.StatusBadRequest).JSON(msg)
	}

	return c.Status(http.StatusOK).JSON(userInfo)
}

func (p *PostApp) HandleDeletePostByID(c *fiber.Ctx) error {
	userInfo, ok := c.UserContext().Value(middleware.CtxUserInfoKey).(types.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusNonAuthoritativeInfo, "user need login")
	}
	return c.JSON(userInfo)
}

func (p *PostApp) GetPostByID(c *fiber.Ctx) error {
	userInfo, ok := c.UserContext().Value(middleware.CtxUserInfoKey).(types.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusNonAuthoritativeInfo, "user need login")
	}
	return c.JSON(userInfo)
}

func (p *PostApp) GetPostsByUserID(c *fiber.Ctx) error {
	userInfo, ok := c.UserContext().Value(middleware.CtxUserInfoKey).(types.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusNonAuthoritativeInfo, "user need login")
	}
	return c.JSON(userInfo)
}

func (p *PostApp) GetPostsByClassfy(c *fiber.Ctx) error {
	userInfo, ok := c.UserContext().Value(middleware.CtxUserInfoKey).(types.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusNonAuthoritativeInfo, "user need login")
	}
	return c.JSON(userInfo)
}
