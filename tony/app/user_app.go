package app

import (
	"fmt"
	"net/http"

	"github.com/HsiaoCz/monster-clone/tony/store"
	"github.com/HsiaoCz/monster-clone/tony/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserApp struct {
	store *store.Store
}

func NewUserApp(store *store.Store) *UserApp {
	return &UserApp{
		store: store,
	}
}

func (u *UserApp) HandleCreateUser(c *fiber.Ctx) error {
	createUserParams := types.CreateUserParams{}
	if err := c.BodyParser(&createUserParams); err != nil {
		return NewAPPError(http.StatusBadRequest, "please check the create user params")
	}
	msg := createUserParams.Validate()
	if len(msg) != 0 {
		return c.Status(http.StatusBadRequest).JSON(msg)
	}
	user := types.NewUserFromParams(createUserParams)
	userRep, err := u.store.US.CreateUser(c.Context(), user)
	if err != nil {
		return NewAPPError(http.StatusInternalServerError, err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "create user success!",
		"user":    userRep,
	})
}

func (u *UserApp) HandleGetUserByID(c *fiber.Ctx) error {
	id := c.Params("uid")
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return NewAPPError(http.StatusBadRequest, "invalid uid")
	}
	user, err := u.store.US.GetUserByID(c.Context(), uid)
	if err != nil {
		return NewAPPError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": http.StatusOK,
		"user":   user,
	})
}

func (u *UserApp) HandleDeleteUserByID(c *fiber.Ctx) error {
	id := c.Params("uid")
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return NewAPPError(http.StatusBadRequest, "invalid uid")
	}
	if err := u.store.US.DeleteUserByID(c.Context(), uid); err != nil {
		return NewAPPError(http.StatusInternalServerError, err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("delete user (uid=%s) success", uid),
	})
}
