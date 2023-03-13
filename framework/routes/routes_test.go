package routes_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eduardor2m/echo-go/framework/routes"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func SetupRouter() *echo.Echo {
	e := echo.New()
	routes.UserRoutes(e)
	return e
}

func TestWelcome(t *testing.T) {
	e := SetupRouter()
	req := httptest.NewRequest(http.MethodGet, "/:name", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:name")
	c.SetParamNames("name")
	c.SetParamValues("Eduardo")
	if assert.NoError(t, routes.WelcomeRoute(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotEqual(t, "Hello Teste", rec.Body.String())
	}
}
