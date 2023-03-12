package controllers

import (
	"net/http"

	"github.com/eduardor2m/echo-go/database"
	"github.com/eduardor2m/echo-go/models"
	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	user := models.User{}
	users := []models.User{}

	if err := c.Bind(&user); err != nil {
		return err
	}

	database.DB.Find(&users)

	for _, u := range users {

		if u.Email == user.Email {
			return c.JSON(http.StatusBadRequest, "Email já cadastrado")
		}

		if u.Username == user.Username {
			return c.JSON(http.StatusBadRequest, "Username já cadastrado")
		}

		if u.CPF == user.CPF {
			return c.JSON(http.StatusBadRequest, "CPF já cadastrado")
		}
	}

	database.DB.Create(&user)

	return c.JSON(http.StatusOK, user)
}

func GetUser(c echo.Context) error {

	email := c.Param("email")
	user := models.User{}

	database.DB.Where("email = ?", email).First(&user)

	if user.Email == "" {
		return c.JSON(http.StatusNotFound, "Usuário não encontrado")
	}

	return c.JSON(http.StatusOK, user)

}

func GetUsers(c echo.Context) error {
	users := []models.User{}

	database.DB.Find(&users)

	if len(users) == 0 {
		return c.JSON(http.StatusNotFound, "Nenhum usuário encontrado")
	}

	return c.JSON(http.StatusOK, users)
}

func UpdateUser(c echo.Context) error {
	user := models.User{}

	email := c.Param("email")

	database.DB.Where("email = ?", email).First(&user)

	if user.Email == "" {
		return c.JSON(http.StatusNotFound, "Usuário não encontrado")
	}

	if err := c.Bind(&user); err != nil {
		return err
	}

	database.DB.Save(&user)

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	user := models.User{}

	email := c.Param("email")

	database.DB.Where("email = ?", email).Delete(&user)

	if user.Email == "" {
		return c.JSON(http.StatusNotFound, "Usuário não encontrado")
	}

	return c.JSON(http.StatusOK, "Usuário deletado com sucesso")
}
