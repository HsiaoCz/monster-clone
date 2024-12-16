package middlewares

import (
	"net/http"

	"github.com/HsiaoCz/monster-clone/lenvenu/storage"
)

type AuthMiddleware struct {
	sen storage.SessionStorer
}

func AuthMiddlewareInit(sen storage.SessionStorer) *AuthMiddleware {
	return &AuthMiddleware{
		sen: sen,
	}
}

func SessionMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
