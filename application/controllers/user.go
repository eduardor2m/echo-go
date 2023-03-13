package controllers

import (
	"net/http"

	"github.com/eduardor2m/echo-go/application/repositories"
	"github.com/eduardor2m/echo-go/application/usecases"
	"github.com/eduardor2m/echo-go/domain"
	"github.com/eduardor2m/echo-go/framework/utils"
	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	user := &domain.User{}
	err := c.Bind(user)

	if err != nil {
		return err
	}

	db := utils.ConnectDb()
	userRepo := repositories.UserRepositoryDB{Db: db}
	useUseCase := usecases.UserUseCase{UserRepository: userRepo}

	user, err = useUseCase.Create(user)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}

func GetUser(c echo.Context) error {
	db := utils.ConnectDb()
	userRepo := repositories.UserRepositoryDB{Db: db}
	useUseCase := usecases.UserUseCase{UserRepository: userRepo}

	id := c.Param("id")

	user, err := useUseCase.Get(id)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func GetUsers(c echo.Context) error {
	db := utils.ConnectDb()
	userRepo := repositories.UserRepositoryDB{Db: db}
	useUseCase := usecases.UserUseCase{UserRepository: userRepo}

	users, err := useUseCase.GetAll()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func UpdateUser(c echo.Context) error {
	user := &domain.User{}
	err := c.Bind(user)

	if err != nil {
		return err
	}

	db := utils.ConnectDb()
	userRepo := repositories.UserRepositoryDB{Db: db}
	useUseCase := usecases.UserUseCase{UserRepository: userRepo}

	user, err = useUseCase.Update(user)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	db := utils.ConnectDb()
	userRepo := repositories.UserRepositoryDB{Db: db}
	useUseCase := usecases.UserUseCase{UserRepository: userRepo}

	id := c.Param("id")

	err := useUseCase.Delete(id)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "User deleted")
}
