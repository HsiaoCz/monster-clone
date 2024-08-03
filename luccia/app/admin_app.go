package app

import (
	"net/http"

	"github.com/HsiaoCz/monster-clone/luccia/store"
)

type AdminApp struct {
	store *store.Store
}

func AdminAppInit(store *store.Store) *AdminApp {
	return &AdminApp{
		store: store,
	}
}

func (a *AdminApp) HandleCreateHotel(w http.ResponseWriter, r *http.Request) error {
	return nil
}
