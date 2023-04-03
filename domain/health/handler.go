package health

import (
	"context"
	"github.com/hellofresh/health-go/v5"
	healthHttp "github.com/hellofresh/health-go/v5/checks/http"
	healthPg "github.com/hellofresh/health-go/v5/checks/postgres"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type Handler struct {
	googleUrl string
	postgres  string
}

func NewHandler(googleUrl, postgres string) *Handler {
	return &Handler{
		googleUrl: googleUrl,
		postgres:  postgres,
	}
}
func (h *Handler) Health(c echo.Context) (err error) {
	check, _ := health.New(health.WithSystemInfo())
	health.WithComponent(health.Component{Name: "library", Version: "1.0.1"})
	err = check.Register(health.Config{
		Name:      "postgres",
		Timeout:   time.Second * 5,
		SkipOnErr: true,
		Check: healthPg.New(healthPg.Config{
			DSN: h.postgres,
		}),
	})
	err = check.Register(health.Config{
		Name:      "google",
		Timeout:   time.Second * 5,
		SkipOnErr: true,
		Check: healthHttp.New(healthHttp.Config{
			URL: h.googleUrl,
		}),
	})
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, check.Measure(context.Background()))
}
