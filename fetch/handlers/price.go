package handlers

import (
	"net/http"
)

type PriceHandler struct{}

func (p *PriceHandler) HandleFetchPrice(w http.ResponseWriter, r *http.Request) error {
	return ErrorMessage(http.StatusInternalServerError, "some thing wrong")
}
