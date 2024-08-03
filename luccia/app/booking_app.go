package app

import (
	"net/http"

	"github.com/HsiaoCz/monster-clone/luccia/store"
)

type BookingApp struct {
	store *store.Store
}

func BookingAppInit(store *store.Store) *BookingApp {
	return &BookingApp{
		store: store,
	}
}

func (b *BookingApp) HandleGetBooking(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (b *BookingApp)HandleUpdateBooking(w http.ResponseWriter, r *http.Request)error{
	return nil
}

func (b *BookingApp)HandleCancelBooking(w http.ResponseWriter, r *http.Request)error{
	return nil
}