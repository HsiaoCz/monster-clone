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
	msg := create_user_params.Validate()
	if len(msg) != 0 {
		return c.Status(http.StatusBadRequest).JSON(msg)
	}
	user := types.NewUserFromParams(create_user_params)
	user, err := u.uc.CreateUser(user)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "create user success",
		"user":    user,
	})
}

func (u *UserHandlers) HandleUserLogin(c *fiber.Ctx) error {
	var user_login_parmas types.UserLoginParmas
	if err := c.BodyParser(&user_login_parmas); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	
	return nil
}
