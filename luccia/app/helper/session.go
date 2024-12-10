package helper

import (
	"context"
	"net/http"

	"github.com/HsiaoCz/monster-clone/luccia/st"
	"github.com/HsiaoCz/monster-clone/luccia/storage"
)

type AuthSession struct {
	sen storage.SessionStorer
}

func AuthSessionInit(sen storage.SessionStorer) *AuthSession {
	return &AuthSession{
		sen: sen,
	}
}

func (a *AuthSession) MiddlewareSession(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//    token:=r.Cookie("token")
		//    there has a problem
		//    should use cache save session token
		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "please login", http.StatusNonAuthoritativeInfo)
			return
		}
		if cookie.Value == "" {
			http.Error(w, "please login", http.StatusNonAuthoritativeInfo)
			return
		}
		session, err := a.sen.GetSessionByToken(r.Context(), cookie.Value)
		if err != nil {
			http.Error(w, "please login", http.StatusNonAuthoritativeInfo)
			return
		}
		ctx := context.WithValue(r.Context(), st.CtxUserSessionKey, session)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
}
