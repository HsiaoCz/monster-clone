package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/HsiaoCz/monster-clone/verson/logger"
)

var StatusCode = &Status{Code: http.StatusOK}

type H map[string]any

func WriteJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	StatusCode.Code = code
	return json.NewEncoder(w).Encode(v)
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func TransferHandlerFunc(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			defer logger.Logger.Error("the http server error", "method", r.Method, "path", r.URL.Path, "remote address", r.RemoteAddr, "error message", err)
			if e, ok := err.(ErrorMsg); ok {
				StatusCode.Code = e.Status
				WriteJSON(w, e.Status, &e)
			} else {
				emsg := ErrorMsg{
					Status:  http.StatusInternalServerError,
					Message: err.Error(),
				}
				StatusCode.Code = emsg.Status
				WriteJSON(w, emsg.Status, &emsg)
			}
		}
		logger.Logger.Info("new request", "method", r.Method, "code", StatusCode.Code, "path", r.URL.Path, "remote address", r.RemoteAddr)
	}
}

type Status struct {
	Code int
}
