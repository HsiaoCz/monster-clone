package handler

import (
	"github.com/HsiaoCz/monster-clone/santino/data"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	user data.UserStorer
}

func NewUserHandler(user data.UserStorer) *UserHandler {
	return &UserHandler{
		user: user,
	}
}

func (u *UserHandler) HandleCreateUser(c *fiber.Ctx) error {
	return nil
}
