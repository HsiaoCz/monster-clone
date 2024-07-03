package db

import (
	"os"

	"github.com/anthdm/superkit/db"
	"github.com/anthdm/superkit/kit"
	"github.com/uptrace/bun"
	_ "github.com/mattn/go-sqlite3"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/extra/bundebug"
)

var Query *bun.DB

func InitDB() error {
	config := db.Config{
		Driver:   os.Getenv("DB_DRIVER"),
		Name:     os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	db, err := db.NewSQL(config)
	if err != nil {
		return err
	}
	Query = bun.NewDB(db, sqlitedialect.New())
	if kit.IsDevelopment() {
		Query.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}
	return nil
}
