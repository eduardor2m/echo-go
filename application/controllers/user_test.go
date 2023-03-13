package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eduardor2m/echo-go/application/controllers"
	"github.com/eduardor2m/echo-go/domain"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func SetupRouter() *echo.Echo {
	e := echo.New()
	e.POST("/users", controllers.CreateUser)
	e.GET("/users/:id", controllers.GetUser)
	e.GET("/users", controllers.GetUsers)
	return e
}

func TestCreateUser(t *testing.T) {
	e := SetupRouter()
	user := domain.User{
		Name:     "Eduardo",
		Email:    "eduardo@gmail.com",
		Password: "123456",
	}
	req := httptest.NewRequest(http.MethodPost, "/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users")
	c.SetParamNames("name", "email", "password")
	c.SetParamValues(user.Name, user.Email, user.Password)
	if assert.NoError(t, controllers.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}

}
