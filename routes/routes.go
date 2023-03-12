package routes

import (
	"github.com/eduardor2m/echo-go/controllers"
	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo) {
	u := e.Group("/api/v1") // Grouping routes
	u.GET("/users", controllers.GetUsers)
	u.GET("/users/:email", controllers.GetUser)
	u.POST("/users", controllers.CreateUser)
	u.PUT("/users/:email", controllers.UpdateUser)
	u.DELETE("/users/:email", controllers.DeleteUser)

}
