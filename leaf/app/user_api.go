package app

import (
	"net/http"

	"github.com/HsiaoCz/monster-clone/leaf/store"
	"github.com/gin-gonic/gin"
)

type UserAPI struct {
	store *store.Store
}

func NewUserAPI(store *store.Store) *UserAPI {
	return &UserAPI{
		store: store,
	}
}

func (u *UserAPI) HandleCreateUser(c *gin.Context) error {
	return NewAPIError(http.StatusInternalServerError, "the server error but we dont know why")
}
