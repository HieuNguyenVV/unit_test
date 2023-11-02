package main

import (
	"github.com/labstack/echo"
	"unit_test/api"
)

func main() {
	e := echo.New()
	e.POST("/create", api.CreateUser)
	e.Start(":8080")
}
