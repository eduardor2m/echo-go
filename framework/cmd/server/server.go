package main

import (
	"github.com/eduardor2m/echo-go/framework/routes"

	"github.com/eduardor2m/echo-go/framework/utils"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	utils.ConnectDb()
	
	routes.UserRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
