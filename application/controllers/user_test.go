package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eduardor2m/echo-go/application/controllers"
	"github.com/eduardor2m/echo-go/domain"
	"github.com/eduardor2m/echo-go/framework/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	e := echo.New()

	reqBody := &domain.User{
		Name:     "Eduardo",
		Email:    "eduardo@gmail.com",
		Password: "123456",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, controllers.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		var user domain.User

		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &user)) {
			assert.Equal(t, "Eduardo", user.Name)
			assert.Equal(t, "eduardo@gmail.com", user.Email)
		}

	}

	defer utils.ConnectDb().Delete(&domain.User{}, "email = ?", reqBody.Email)

}

func TestListAllUsers(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, controllers.GetUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

}

func TestGetUser(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/users/2", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, controllers.GetUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

}

func TestUpdateUser(t *testing.T) {
	e := echo.New()

	reqBody := &domain.User{
		Name:     "Eduardo Update",
		Email:    "eduardoupdate@gmail.com",
		Password: "123456",
	}

	reqJSON, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPut, "/users/2", bytes.NewBuffer(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, controllers.UpdateUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

}

func TestDeleteUser(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodDelete, "/users/2", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, controllers.DeleteUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

}
