package app

import (
	"encoding/json"
	"net/http"

	"github.com/HsiaoCz/monster-clone/luccia/st"
	"github.com/HsiaoCz/monster-clone/luccia/store"
)

type UserApp struct {
	user store.UserStorer
}

func UserAppInit(user store.UserStorer) *UserApp {
	return &UserApp{
		user: user,
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
	result, err := u.user.CreateUser(r.Context(), user)
	if err != nil {
		return ErrorMessage(http.StatusInternalServerError, err.Error())
	}
	return WriteJson(w, http.StatusOK, map[string]any{
		"message": "create user success!",
		"data":    result,
	})
}
