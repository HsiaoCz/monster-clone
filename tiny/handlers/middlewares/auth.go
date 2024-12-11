package middlewares

import (
	"net/http"

	"github.com/HsiaoCz/monster-clone/luna/handlers"
	"github.com/HsiaoCz/monster-clone/luna/storage"
	"github.com/HsiaoCz/monster-clone/luna/types"
	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	sen storage.SessionStoreInter
}

func AuthMiddlewareInit(sen storage.SessionStoreInter) *AuthMiddleware {
	return &AuthMiddleware{
		sen: sen,
	}
}

func (a *AuthMiddleware) MiddlewareAuth(c *fiber.Ctx) error {
	token := c.Cookies("token")
	if token == "" {
		return handlers.ErrorMessage(http.StatusNonAuthoritativeInfo, "please login")
	}
	// there should use redis to get token
	session, err := a.sen.GetSessionByID(c.Context(), token)
	if err != nil {
		return handlers.ErrorMessage(http.StatusNonAuthoritativeInfo, "please login")
	}
	c.Locals(types.CtxUserInfo, session)
	return c.Next()
}
