package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/HsiaoCz/monster-clone/lenvenu/types"
)

type CtxKey string

const (
	CtxUserInfoKey CtxKey = "userInfo"
)

// JWT Middleware
func JWTAuthMiddleware(netx http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing authorization header", http.StatusUnauthorized)
			return
		}

		// Split the header to get the token part
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 ||
			parts[0] != "Bearer" {
			http.Error(w, "invalid authorization header format", http.StatusUnauthorized)
			return
		}
		tokenString := parts[1]

		// parse and validate the token
		mc, err := ParseToken(tokenString)
		if err != nil {
			http.Error(w, "invalid token string", http.StatusUnauthorized)
			return
		}
		userInfo := &types.UserInfo{
			UserID:  mc.UserID,
			Email:   mc.Email,
			IsAdmin: mc.IsAdmin,
		}
		ctx := context.WithValue(r.Context(), CtxUserInfoKey, userInfo)
		r = r.WithContext(ctx)

		// Pass the request to the next handler
		netx.ServeHTTP(w, r)
	})
}
