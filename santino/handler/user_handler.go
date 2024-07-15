package handler

import (
	"log"
	"net/http"

	"github.com/HsiaoCz/monster-clone/santino/data"
	"github.com/HsiaoCz/monster-clone/santino/types"
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
	var userCreateParams types.CreateUserParams
	if err := c.BodyParser(&userCreateParams); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	msg := userCreateParams.Validate()
	if len(msg) != 0 {
		return c.Status(http.StatusBadRequest).JSON(msg)
	}
	user := types.NewUserFromParams(userCreateParams)
	result, err := u.user.CreateUser(user)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "create user success",
		"user":    result,
	})
}

func (u *UserHandler) HandleGetUserByID(c *fiber.Ctx) error {
	user_id := c.Params("user_id")
	log.Println(user_id)
	user, err := u.user.GetUserByID(user_id)
	if err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "get user success",
		"user":    user,
	})
}
