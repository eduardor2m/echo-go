package routes

import (
	"github.com/eduardor2m/echo-go/application/controllers"
	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
	e.POST("/users", controllers.CreateUser)
}
