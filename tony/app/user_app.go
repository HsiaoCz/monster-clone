package app

import (
	"fmt"
	"net/http"

	"github.com/HsiaoCz/monster-clone/tony/app/middleware"
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
		return ErrorMessage(http.StatusBadRequest, "please check the create user params")
	}
	msg := createUserParams.Validate()
	if len(msg) != 0 {
		return c.Status(http.StatusBadRequest).JSON(msg)
	}
	user := types.NewUserFromParams(createUserParams)
	userRep, err := u.store.US.CreateUser(c.Context(), user)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
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
		return ErrorMessage(http.StatusBadRequest, "invalid uid")
	}
	user, err := u.store.US.GetUserByID(c.Context(), uid)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": http.StatusOK,
		"user":   user,
	})
}

func (u *UserApp) HandleDeleteUserByID(c *fiber.Ctx) error {
	// delete user need user login
	// so the userID should get in the user context
	// id := c.Params("uid")
	// uid, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	return ErrorMessage(http.StatusBadRequest, "invalid uid")
	// }
	userInfo, ok := c.UserContext().Value(middleware.CtxUserInfoKey).(*types.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusNonAuthoritativeInfo, "user need login")
	}
	if err := u.store.US.DeleteUserByID(c.Context(), userInfo.UserID); err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("delete user (uid=%s) success", userInfo.UserID),
	})
}

func (u *UserApp) HandleUpdateUserByID(c *fiber.Ctx) error {
	// update user need user login
	// so the userID should get in the user context
	userInfo, ok := c.UserContext().Value(middleware.CtxUserInfoKey).(*types.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusNonAuthoritativeInfo, "user need login")
	}
	userUpdateParams := types.UpdateUserParmas{}
	if err := c.BodyParser(&userUpdateParams); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	if msg := userUpdateParams.Validate(); len(msg) != 0 {
		return c.Status(http.StatusBadRequest).JSON(msg)
	}
	userInsertDBUpdateParams := types.NewInstertDBUpdateUserParams(userUpdateParams)
	user, err := u.store.US.UpdateUserByID(c.Context(), userInfo.UserID, userInsertDBUpdateParams)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": http.StatusOK,
		"user":   user,
	})
}
