package app

import (
	"net/http"

	"github.com/HsiaoCz/monster-clone/luccia/storage"
)

type HotelApp struct {
	store *storage.Store
}

func HotelAppInit(store *storage.Store) *HotelApp {
	return &HotelApp{
		store: store,
	}
}

func (h *HotelApp) HandleGetHotels(w http.ResponseWriter, r *http.Request) error {
	// get hotels don't need login
	return nil
}

func (h *HotelApp) HandleGetHotelByID(w http.ResponseWriter, r *http.Request) error {
	// get hotels don't need login
	return nil
}
