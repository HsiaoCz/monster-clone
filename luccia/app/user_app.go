package app

import (
	"encoding/json"
	"net/http"

	"github.com/HsiaoCz/monster-clone/luccia/app/helper"
	"github.com/HsiaoCz/monster-clone/luccia/st"
	"github.com/HsiaoCz/monster-clone/luccia/store"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserApp struct {
	store *store.Store
}

func UserAppInit(store *store.Store) *UserApp {
	return &UserApp{
		store: store,
	}
}

func (u *UserApp) HandleCreateUser(w http.ResponseWriter, r *http.Request) error {
	var user_create_params st.CreateUserParam
	if err := json.NewDecoder(r.Body).Decode(&user_create_params); err != nil {
		return ErrorMessage(http.StatusBadRequest, "please check the request params")
	}
	msg := user_create_params.ValidateCreateUserParam()
	if len(msg) != 0 {
		return WriteJson(w, http.StatusBadRequest, msg)
	}
	user := st.NewUserFromReq(user_create_params)
	result, err := u.store.Us.CreateUser(r.Context(), user)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJson(w, http.StatusOK, map[string]any{
		"message": "create user success!",
		"data":    result,
	})
}

func (u *UserApp) HandleUserLogin(w http.ResponseWriter, r *http.Request) error {
	var user_login_params st.UserLoginParams
	if err := json.NewDecoder(r.Body).Decode(&user_login_params); err != nil {
		return ErrorMessage(http.StatusBadRequest, "please check the email or password")
	}
	parmas := user_login_params.EncryptedPassword()
	user, err := u.store.Us.GetUserByEmail(r.Context(), parmas.Email)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	if parmas.Password != user.Password {
		return ErrorMessage(http.StatusBadRequest, "please check the email or passwrod")
	}
	token, err := helper.GenToken(user.ID, user.Email, user.IsAdmin)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJson(w, http.StatusOK, map[string]any{
		"status": http.StatusOK,
		"token":  token,
		"user":   user,
	})
}

func (u *UserApp) HandleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	id := r.URL.Query().Get("uid")
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrorMessage(http.StatusBadRequest, "query param invalid")
	}
	user, err := u.store.Us.GetUserByID(r.Context(), uid)
	if err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	return WriteJson(w, http.StatusOK, user)
}

func (u *UserApp) HandleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	userInfo, ok := r.Context().Value(st.CtxUserInfoKey).(*st.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusUnauthorized, "please login")
	}
	if err := u.store.Us.DeleteUserByID(r.Context(), userInfo.UserID); err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJson(w, http.StatusOK, map[string]any{
		"status":  http.StatusOK,
		"message": "delete user success",
	})
}

func (u *UserApp) HandleUpdateUser(w http.ResponseWriter, r *http.Request) error {
	userInfo, ok := r.Context().Value(st.CtxUserInfoKey).(*st.UserInfo)
	if !ok {
		return ErrorMessage(http.StatusUnauthorized, "please login")
	}
	var up st.UpdateUserParams
	if err := json.NewDecoder(r.Body).Decode(&up); err != nil {
		return ErrorMessage(http.StatusBadRequest, "please check the update params")
	}
	msg := up.ValidateUpdateUserParams()
	if len(msg) != 0 {
		return WriteJson(w, http.StatusBadRequest, msg)
	}
	user, err := u.store.Us.UpdateUser(r.Context(), userInfo.UserID, &up)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJson(w, http.StatusOK, map[string]any{
		"status": http.StatusOK,
		"user":   user,
	})
}

func (u *UserApp) HandleUserVerifyPassword(w http.ResponseWriter, r *http.Request) error {
	userInfo, ok := r.Context().Value(st.CtxUserInfoKey).(*st.User)
	if !ok {
		return ErrorMessage(http.StatusUnauthorized, "please login")
	}
	var user_verify_passwd_params st.VerifyUserPasswordParmas
	if err := json.NewDecoder(r.Body).Decode(&user_verify_passwd_params); err != nil {
		return ErrorMessage(http.StatusBadRequest, err.Error())
	}
	if !user_verify_passwd_params.Validate() {
		return ErrorMessage(http.StatusBadRequest, "verify your password")
	}
	if err := u.store.Us.VerifyUserPassword(r.Context(), userInfo.ID, user_verify_passwd_params.EncryptedUserPassword()); err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJson(w, http.StatusOK, map[string]any{
		"status":  http.StatusOK,
		"message": "chenge password success",
	})
}

func (u *UserApp) HandleUserBookingRoom(w http.ResponseWriter, r *http.Request) error {
	return nil
}
