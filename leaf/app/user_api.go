package app

import (
	"fmt"
	"net/http"

	"github.com/HsiaoCz/monster-clone/leaf/app/middlewares"
	"github.com/HsiaoCz/monster-clone/leaf/models"
	"github.com/HsiaoCz/monster-clone/leaf/store"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		return ErrorMessage(http.StatusBadRequest, "please check the create user params")
	}
	msg := createUserParams.Validate()
	if len(msg) != 0 {
		return c.Status(http.StatusBadRequest).JSON(msg)
	}
	user := models.NewUserFromParams(createUserParams)
	userRep, err := u.store.User.CreateUser(c.Context(), user)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "create user success!",
		"user":    userRep,
	})
}

func (u *UserAPI) HandleGetUserByID(c *fiber.Ctx) error {
	id := c.Params("uid")
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrorMessage(http.StatusBadRequest, "invalid uid")
	}
	user, err := u.store.User.GetUserByID(c.Context(), uid)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": http.StatusOK,
		"user":   user,
	})
}

func (u *UserAPI) HandleDeleteUserByID(c *fiber.Ctx) error {
	id := c.Params("uid")
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrorMessage(http.StatusBadRequest, "invalid uid")
	}
	if err := u.store.User.DeleteUserByID(c.Context(), uid); err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("delete user (uid=%s) success", uid),
	})
}

func (u *UserAPI) HandleUpdateUser(c *fiber.Ctx) error {
	updateUser := models.UpdateUserParmas{}
	if err := c.BodyParser(&updateUser); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	msg := updateUser.Validate()
	if len(msg) != 0 {
		return c.Status(http.StatusBadRequest).JSON(msg)
	}
	userInfo, ok := c.UserContext().Value(middlewares.CtxUserInfoKey).(*models.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusNonAuthoritativeInfo, "user need login")
	}
	user, err := u.store.User.UpdateUserByID(c.Context(), userInfo.UserID, &updateUser)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": http.StatusOK,
		"user":   user,
	})
}

func (u *UserAPI) HandleUpdatePassword(c *fiber.Ctx) error {
	email := c.Params("email")

	userInfo, ok := c.UserContext().Value(middlewares.CtxUserInfoKey).(models.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusNonAuthoritativeInfo, "need login")
	}

	if email != userInfo.Email {
		return ErrorMessage(http.StatusBadRequest, "please check the email")
	}

	// there hava a problem
	// verify the email and reset the password should use two handlers

	updateUserPasswd := models.UpdateUserPassword{}

	if err := c.BodyParser(&updateUserPasswd); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}

	if err := updateUserPasswd.Validate(); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}

	if err := u.store.User.UpdateUserPassword(c.Context(), userInfo.UserID, models.NewPasswordFromParam(updateUserPasswd)); err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "change the password success",
	})
}
