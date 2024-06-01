package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/HsiaoCz/monster-clone/leaf/app"
	v1 "github.com/HsiaoCz/monster-clone/leaf/app/v1"
	"github.com/HsiaoCz/monster-clone/leaf/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	var (
		user     = os.Getenv("MUSER")
		password = os.Getenv("PASSWORD")
		host     = os.Getenv("HOST")
		port     = os.Getenv("PORT")
		dbname   = os.Getenv("DBNAME")
	)
	// we don't need this shit
	// but if the db connect error we use this to check shit where?
	// fmt.Println(user, password, host, port, dbname)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	var (
		logger         = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
		mysqlUserStore = store.NewMySqlUserStore(db)
		store          = &store.Store{User: mysqlUserStore}
		userHandlers   = v1.NewUserAPI(store)
		router         = gin.Default()
		v1             = router.Group("/app/v1")
	)
	slog.SetDefault(logger)
	// routers
	{
		v1.POST("/user", app.TransferHandlerfunc(userHandlers.HandleCreateUser))
	}
	router.Run(os.Getenv("LISTENADDR"))
}
