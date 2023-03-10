package routes

import (
	"github.com/eduardor2m/echo-go/controllers"
	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	e.GET("/", controllers.CreateUser)
}