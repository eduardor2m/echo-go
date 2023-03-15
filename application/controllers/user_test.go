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

	reqBody := map[string]string{
		"name":     "Eduardo",
		"email":    "eduardo@gmail.com",
		"password": "123456",
	}

	defer func() {
		db := utils.ConnectDb()
		db.Exec("DROP DATABASE users")
	}()

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

}
