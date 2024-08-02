package app

import "github.com/HsiaoCz/monster-clone/luccia/store"

type BookingApp struct {
	store *store.Store
}

func BookingAppInit(store *store.Store)*BookingApp{
	return &BookingApp{
		store: store,
	}
}

