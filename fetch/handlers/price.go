package handlers

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func HandleFetchPrice(w http.ResponseWriter, r *http.Request) {
	logrus.WithFields(logrus.Fields{
		"message": "hello",
	}).Info()
}
