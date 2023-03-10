package main

import (
	"github.com/eduardor2m/echo-go/routes"
	"github.com/labstack/echo/v4"
)

func main() {
  e := echo.New()

  routes.UserRoute(e)

  e.Logger.Fatal(e.Start(":1323"))
}
