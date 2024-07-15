package handler

import (
	"github.com/HsiaoCz/monster-clone/santino/data"
	"github.com/gofiber/fiber/v2"
)

type AdminHandler struct {
	admin data.AdminStorer
}

func NewAdminHandler(admin data.AdminStorer) *AdminHandler {
	return &AdminHandler{
		admin: admin,
	}
}

func (a *AdminHandler) HandleCreateAdmin(c *fiber.Ctx) error {
	return nil
}
