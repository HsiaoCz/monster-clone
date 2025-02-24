package main

import (
	"net/http"
	"os"

	"github.com/HsiaoCz/monster-clone/cloud/logger"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logger.L.Error("load env error", err)
		os.Exit(1)
	}

	var (
		port = os.Getenv("PORT")
		app  = http.NewServeMux()
	)

	app.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	logger.L.WithFields(logrus.Fields{
		"port": port,
	}).Info("server start")
	if err := http.ListenAndServe(port, app); err != nil {
		logger.L.Error("server start error", err)
		os.Exit(1)
	}

}
