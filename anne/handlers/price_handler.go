package handlers

import "net/http"

type PriceHandler struct{}

func (s *PriceHandler) HandlePriceFetch(w http.ResponseWriter, r *http.Request) error {
	return ErrorMessage(http.StatusInternalServerError, "something wrong")
}
