package handlers

import (
	"net/http"

	"github.com/HsiaoCz/monster-clone/peek/services"
	"github.com/HsiaoCz/monster-clone/peek/types"
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
	var create_user_params types.CreateUserParams
	if err := c.BodyParser(create_user_params); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	
	return nil
}
