package handlers

import "github.com/gofiber/fiber/v2"

type UserApp struct{}

func (u *UserApp) HandleCreateUser(c *fiber.Ctx) error {
	return nil
}
