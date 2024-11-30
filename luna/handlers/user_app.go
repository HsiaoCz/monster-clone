package handlers

import (
	"github.com/HsiaoCz/monster-clone/luna/storage"
	"github.com/gofiber/fiber/v2"
)

type UserApp struct {
	store   storage.UserStoreInter
	session storage.SessionStoreInter
}

func UserAppInit(store storage.UserStoreInter, session storage.SessionStoreInter) *UserApp {
	return &UserApp{
		store:   store,
		session: session,
	}
}

func (u *UserApp) HandleCreateUser(c *fiber.Ctx) error {
	return nil
}
