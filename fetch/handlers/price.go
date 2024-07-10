package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/HsiaoCz/monster-clone/fetch/services"
	"github.com/sirupsen/logrus"
)

type PriceHandler struct {
	listenAddr string
	svc        services.FetchPricer
}

func NewPriceHandler(listenAddr string, svc services.FetchPricer) *PriceHandler {
	return &PriceHandler{
		listenAddr: listenAddr,
		svc:        svc,
	}
}

func (p *PriceHandler) HandleFetchPrice(w http.ResponseWriter, r *http.Request) error {
	return ErrorMessage(http.StatusInternalServerError, "some thing wrong")
}

func (p *PriceHandler) Run() {
	router := http.NewServeMux()

	router.HandleFunc("GET /price", TransferHandlerfunc(p.HandleFetchPrice))

	logrus.WithFields(logrus.Fields{
		"listen address": p.listenAddr,
	}).Info("the http server is running")
	http.ListenAndServe(p.listenAddr, router)
}

func writeJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}
