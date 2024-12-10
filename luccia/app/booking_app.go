package app

import (
	"net/http"

	"github.com/HsiaoCz/monster-clone/luccia/st"
	"github.com/HsiaoCz/monster-clone/luccia/storage"
)

type BookingApp struct {
	store *storage.Store
}

func BookingAppInit(store *storage.Store) *BookingApp {
	return &BookingApp{
		store: store,
	}
}

func (b *BookingApp) HandleGetBooking(w http.ResponseWriter, r *http.Request) error {
	// get bookings need login
	userInfo, ok := r.Context().Value(st.CtxUserInfoKey).(*st.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusUnauthorized, "user unlogin")
	}
	if !userInfo.IsAdmin {
		return ErrorMessage(http.StatusUnauthorized, "can't do this shit")
	}
	return nil
}

func (b *BookingApp) HandleUpdateBooking(w http.ResponseWriter, r *http.Request) error {
	userInfo, ok := r.Context().Value(st.CtxUserInfoKey).(*st.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusUnauthorized, "user unlogin")
	}
	if !userInfo.IsAdmin {
		return ErrorMessage(http.StatusUnauthorized, "can't do this shit")
	}
	return nil
}

func (b *BookingApp) HandleCancelBooking(w http.ResponseWriter, r *http.Request) error {
	userInfo, ok := r.Context().Value(st.CtxUserInfoKey).(*st.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusUnauthorized, "user unlogin")
	}
	if !userInfo.IsAdmin {
		return ErrorMessage(http.StatusUnauthorized, "can't do this shit")
	}
	return nil
}
