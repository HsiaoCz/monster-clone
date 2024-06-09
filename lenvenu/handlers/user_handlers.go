package handlers

import (
	"net/http"

	"github.com/HsiaoCz/monster-clone/lenvenu/db"
)

type UserHandlers struct {
	store *db.Store
}

func NewUserHandlers(store *db.Store) *UserHandlers {
	return &UserHandlers{
		store: store,
	}
}

func (u *UserHandlers) HandleCreateUser(w http.ResponseWriter, r *http.Request) error {
	return WriteJSON(w, http.StatusOK, "hello everyone")
}
