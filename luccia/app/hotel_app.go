package app

import (
	"net/http"

	"github.com/HsiaoCz/monster-clone/luccia/store"
)

type HotelApp struct {
	store *store.Store
}

func HotelAppInit(store *store.Store) *HotelApp {
	return &HotelApp{
		store: store,
	}
}

func (h *HotelApp) HandleGetHotels(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *HotelApp) HandleGetHotelByID(w http.ResponseWriter, r *http.Request) error {
	return nil
}
