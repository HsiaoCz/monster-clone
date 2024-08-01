package app

import "github.com/HsiaoCz/monster-clone/luccia/store"

type HotelApp struct {
	store *store.Store
}

func HotelAppInit(store *store.Store) *HotelApp {
	return &HotelApp{
		store: store,
	}
}
