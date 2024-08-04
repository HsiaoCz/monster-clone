package app

import (
	"net/http"

	"github.com/HsiaoCz/monster-clone/luccia/st"
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
	userInfo, ok := r.Context().Value(st.CtxUserInfoKey).(*st.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusUnauthorized, "user unlogin")
	}
	if !userInfo.IsAdmin {
		return ErrorMessage(http.StatusUnauthorized, "can't do this shit")
	}
	return nil
}

func (a *AdminApp) HandleCreateRoom(w http.ResponseWriter, r *http.Request) error {
	userInfo, ok := r.Context().Value(st.CtxUserInfoKey).(*st.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusUnauthorized, "user unlogin")
	}
	if !userInfo.IsAdmin {
		return ErrorMessage(http.StatusUnauthorized, "can't do this shit")
	}
	return nil
}

func (a *AdminApp) HandleDeleteHotel(w http.ResponseWriter, r *http.Request) error {
	userInfo, ok := r.Context().Value(st.CtxUserInfoKey).(*st.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusUnauthorized, "user unlogin")
	}
	if !userInfo.IsAdmin {
		return ErrorMessage(http.StatusUnauthorized, "can't do this shit")
	}
	return nil
}

func (a *AdminApp) HandleDeleteRoom(w http.ResponseWriter, r *http.Request) error {
	userInfo, ok := r.Context().Value(st.CtxUserInfoKey).(*st.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusUnauthorized, "user unlogin")
	}
	if !userInfo.IsAdmin {
		return ErrorMessage(http.StatusUnauthorized, "can't do this shit")
	}
	return nil
}
