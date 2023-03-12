package main

import (
	"github.com/eduardor2m/echo-go/database"
	"github.com/eduardor2m/echo-go/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	database.InitDB()

	routes.Route(e)

	e.Logger.Fatal(e.Start(":1323"))
}
