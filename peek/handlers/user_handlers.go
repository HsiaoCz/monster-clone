package handlers

import (
	"github.com/HsiaoCz/monster-clone/peek/services"
	"github.com/gofiber/fiber/v2"
)

type UserHandlers struct {
	uc services.UserCaseInter
}

func NewUserHandlers(uc services.UserCaseInter) *UserHandlers {
	return &UserHandlers{
		uc: uc,
	}
}

func (u *UserHandlers) HandleCreateUser(c *fiber.Ctx) error {
	return nil
}
