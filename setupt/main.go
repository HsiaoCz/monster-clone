package main

import (
	"github.com/HsiaoCz/monster-clone/setupt/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	// app.GET("/hello", func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, echo.Map{
	// 		"message": "all is well",
	// 	})
	// })
	userHandler := &handler.UserHandler{}
	app.GET("/user", userHandler.HandleUserShow)
	
	app.Start(":3001")
}
