package handler

import (
	"net/http"

	"github.com/HsiaoCz/monster-clone/setupt/model"
	"github.com/HsiaoCz/monster-clone/setupt/view/userv"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h *UserHandler) HandleUserShow(c echo.Context) error {
	user := model.User{
		Username: "bob",
		Email:    "asd@gmail.com",
	}
	return Render(c, http.StatusOK, userv.Show(user))
}
