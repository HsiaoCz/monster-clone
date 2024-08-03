package app

import (
	"net/http"

	"github.com/HsiaoCz/monster-clone/luccia/store"
)

type RoomApp struct {
	store *store.Store
}

func RoomAppInit(store *store.Store) *RoomApp {
	return &RoomApp{
		store: store,
	}
}

func (ra *RoomApp) HandleGetRooms(w http.ResponseWriter, r *http.Request) error {
	return nil
}
