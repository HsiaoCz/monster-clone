package handlers

import (
	"context"
	"encoding/json"
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/HsiaoCz/monster-clone/fetch/types"
	"github.com/sirupsen/logrus"
)

var StatusCode = &Status{Code: http.StatusOK}

type Handlerfunc func(w http.ResponseWriter, r *http.Request) error

type Status struct {
	Code int
}

func WriteJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Set("Content-Type", "application")
	w.WriteHeader(code)
	StatusCode.Code = code
	return json.NewEncoder(w).Encode(v)
}

func TransferHandlerfunc(h Handlerfunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestID := rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(math.MaxInt64)
		ctx := context.WithValue(r.Context(), types.CtxRequestIDKey, requestID)
		r = r.WithContext(ctx)
		start := time.Now()
		if err := h(w, r); err != nil {
			defer func() {
				logrus.WithFields(logrus.Fields{
					"requestID":      requestID,
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
				errMsg := ErrorMsg{
					Status:  http.StatusInternalServerError,
					Message: err.Error(),
				}
				StatusCode.Code = errMsg.Status
				WriteJSON(w, errMsg.Status, &errMsg)
			}
		}
		logrus.WithFields(logrus.Fields{
			"requestID":      requestID,
			"method":         r.Method,
			"code":           StatusCode.Code,
			"path":           r.URL.Path,
			"remote address": r.RemoteAddr,
			"cost":           time.Since(start),
		}).Info("new request comming")
	}
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
