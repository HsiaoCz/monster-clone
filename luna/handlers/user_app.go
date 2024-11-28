package handlers

import (
	"github.com/HsiaoCz/monster-clone/luna/storage"
	"github.com/gofiber/fiber/v2"
)

type UserApp struct {
	store storage.UserStoreInter
}

func UserAppInit(store storage.UserStoreInter) *UserApp {
	return &UserApp{
		store: store,
	}
}

func (u *UserApp) HandleCreateUser(c *fiber.Ctx) error {
	return nil
}
