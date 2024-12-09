package handlers

import (
	"net/http"
	"time"

	"github.com/HsiaoCz/monster-clone/luna/storage"
	"github.com/HsiaoCz/monster-clone/luna/types"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserApp struct {
	store   storage.UserStoreInter
	session storage.SessionStoreInter
}

func UserAppInit(store storage.UserStoreInter, session storage.SessionStoreInter) *UserApp {
	return &UserApp{
		store:   store,
		session: session,
	}
}

func (u *UserApp) HandleCreateUser(c *fiber.Ctx) error {
	var create_user_params types.CreateUserParams
	if err := c.BodyParser(&create_user_params); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	msg := create_user_params.Validate()
	if len(msg) != 0 {
		return c.Status(http.StatusBadRequest).JSON(msg)
	}
	user := types.NewUserFromParams(create_user_params)
	user, err := u.store.CreateUser(c.Context(), user)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	// session should create after user login
	session := &types.Sessions{
		Token:     uuid.New().String(),
		UserID:    user.ID.String(),
		IpAddress: c.IP(),
		UserAgent: string(c.Request().Header.UserAgent()),
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30),
	}
	session, err = u.session.CreateSession(c.Context(), session)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    session.Token,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		HTTPOnly: true,
	})
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "create user success",
		"user":    user,
		"session": session,
	})
}

func (u *UserApp)HandleUpdateUser(c *fiber.Ctx)error{
	return c.Status(http.StatusOK).JSON(fiber.Map{
		
	})
}