package app

import (
	"net/http"

	"github.com/HsiaoCz/monster-clone/luccia/storage"
)

type RoomApp struct {
	store *storage.Store
}

func RoomAppInit(store *storage.Store) *RoomApp {
	return &RoomApp{
		store: store,
	}
}

func (ra *RoomApp) HandleGetRooms(w http.ResponseWriter, r *http.Request) error {
	// get hotels don't need login
	return nil
}

func (ra *RoomApp) HandleGetRoomByID(w http.ResponseWriter, r *http.Request) error {
	// get hotels don't need login
	return nil
}

func (ra *RoomApp) HandleBookingRoom(w http.ResponseWriter, r *http.Request) error {
	return nil
}
