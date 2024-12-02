package middlewares

import (
	"errors"

	"github.com/HsiaoCz/monster-clone/peek/services"
	"github.com/HsiaoCz/monster-clone/peek/types"
	"github.com/gofiber/fiber/v2"
)

type AuthMiddle struct {
	sn services.SessionCaser
}

func AuthMiddleInit(sn services.SessionCaser) *AuthMiddle {
	return &AuthMiddle{
		sn: sn,
	}
}

func (a *AuthMiddle) AuthMiddleware(c *fiber.Ctx) error {
	token := c.Cookies("token")
	if token == "" {
		return errors.New("please login")
	}
	session, err := a.sn.GetSessionByToken(c.Context(), token)
	if err != nil {
		return errors.New("please login")
	}
	c.Locals(types.CtxUserInfo, session)
	return c.Next()
}
