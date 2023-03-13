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
