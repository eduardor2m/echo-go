package main

import (
	"github.com/eduardor2m/echo-go/framework/routes"

	"github.com/eduardor2m/echo-go/framework/utils"
	"github.com/labstack/echo/v4"
)

func main() {
	utils.ConnectDb()

	// user := &domain.User{
	// 	Name:     "Eduardo",
	// 	Email:    "eduardo@gmail.com",
	// 	Password: "123456",
	// }

	// userRepo := repositories.UserRepositoryDB{Db: db}
	// userUsecase := usecases.UserUseCase{UserRepository: userRepo}
	// result, err := userUsecase.Create(user)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(result.Name, result.Email, result.Password)

	e := echo.New()

	routes.UserRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))

}
