package logger

import (
	"io"
	"log/slog"
	"os"
)

var Logger *logger

type logger struct {
	infofile  io.Writer
	debugfile io.Writer
	errorfile io.Writer
	warnfile  io.Writer
	log       *slog.Logger
}

func InitLogger() error {
	logger := &logger{}
	infofile, err := os.OpenFile("./info.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	logger.infofile = infofile
	errfile, err := os.OpenFile("./error.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	logger.errorfile = errfile
	debugfile, err := os.OpenFile("./debug.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	logger.debugfile = debugfile
	warnfile, err := os.OpenFile("./warn.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	logger.warnfile = warnfile
	Logger = logger
	return nil
}
func (l *logger) Info(msg string, args ...any) {
	l.log = slog.New(slog.NewJSONHandler(io.MultiWriter(l.infofile, os.Stdout), &slog.HandlerOptions{}))
	l.log.Info(msg, args...)
}

func (l *logger) Debug(msg string, args ...any) {
	l.log = slog.New(slog.NewJSONHandler(io.MultiWriter(l.debugfile, os.Stdout), &slog.HandlerOptions{}))
	l.log.Debug(msg, args...)
}

func (l *logger) Error(msg string, args ...any) {
	l.log = slog.New(slog.NewJSONHandler(io.MultiWriter(l.errorfile, os.Stdout), &slog.HandlerOptions{}))
	l.log.Error(msg, args...)
}

func (l *logger) Warn(msg string, args ...any) {
	l.log = slog.New(slog.NewJSONHandler(io.MultiWriter(l.warnfile, os.Stdout), &slog.HandlerOptions{}))
	l.log.Warn(msg, args...)
}
