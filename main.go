package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	r "lab4-kelsingazin/utils"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/users", r.CreateUser)
	e.GET("/users/:id", r.GetUser)
	e.GET("/users", r.GetAllUsers)
	e.DELETE("/users/:id", r.DeleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}
