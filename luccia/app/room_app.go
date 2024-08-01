package app

import "github.com/HsiaoCz/monster-clone/luccia/store"

type RoomApp struct {
	store *store.Store
}

func RoomAppInit(store *store.Store) *RoomApp {
	return &RoomApp{
		store: store,
	}
}
