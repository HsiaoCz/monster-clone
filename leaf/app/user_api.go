package app

import (
	"net/http"

	"github.com/HsiaoCz/monster-clone/leaf/models"
	"github.com/HsiaoCz/monster-clone/leaf/store"
	"github.com/gofiber/fiber/v2"
)

type UserAPI struct {
	store *store.Store
}

func NewUserAPI(store *store.Store) *UserAPI {
	return &UserAPI{
		store: store,
	}
}

func (u *UserAPI) HandleCreateUser(c *fiber.Ctx) error {
	createUserParams := models.CreateUserParams{}
	if err := c.BodyParser(&createUserParams); err != nil {
		return NewAPIError(http.StatusBadRequest, "please check the create user params")
	}
	msg := createUserParams.Validate()
	if len(msg) != 0 {
		return c.Status(http.StatusBadRequest).JSON(msg)
	}
	user := models.NewUserFromParams(createUserParams)
	userRep, err := u.store.User.CreateUser(c.Context(), user)
	if err != nil {
		return NewAPIError(http.StatusInternalServerError, err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "create user success!",
		"user":    userRep,
	})
}

func (u *UserAPI) HandleGetUserByID(c *fiber.Ctx) error {
	return nil
}
