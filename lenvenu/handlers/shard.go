package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

var StatusCode = &Status{Code: http.StatusOK}

type H map[string]any

func WriteJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	StatusCode.Code = code
	return json.NewEncoder(w).Encode(v)
}

type ErrorMsg struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e ErrorMsg) Error() string {
	return e.Message
}

func ErrorMessage(status int, message string) ErrorMsg {
	return ErrorMsg{
		Status:  status,
		Message: message,
	}
}

type Status struct {
	Code int
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func TransferHandlerFunc(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		if err := h(w, r); err != nil {
			defer func() {
				logrus.WithFields(logrus.Fields{
					"method":         r.Method,
					"path":           r.URL.Path,
					"remote address": r.RemoteAddr,
					"error message":  err,
				}).Error("the http server error")
			}()
			if e, ok := err.(ErrorMsg); ok {
				StatusCode.Code = e.Status
				WriteJSON(w, e.Status, &e)
			} else {
				arr := ErrorMsg{
					Status:  http.StatusInternalServerError,
					Message: err.Error(),
				}
				StatusCode.Code = arr.Status
				WriteJSON(w, arr.Status, &arr)
			}
		}
		logrus.WithFields(logrus.Fields{
			"method":         r.Method,
			"code":           StatusCode.Code,
			"path":           r.URL.Path,
			"remote address": r.RemoteAddr,
			"cost":           time.Since(start),
		}).Info("new request comming")
		slog.Info("new request", "method", r.Method, "code", status.Code, "path", r.URL.Path, "remote address", r.RemoteAddr)
	}
}
