package internal

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.opentelemetry.io/otel"
)

type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	users   = map[int]*user{}
	seq     = 1
	counter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "gss_single",
	})
)

const name = "gss"

func CreateUser(c echo.Context) error {
	_, span := otel.Tracer(name).Start(c.Request().Context(), "Create")
	defer span.End()
	u := &user{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func GetUser(c echo.Context) error {
	_, span := otel.Tracer(name).Start(c.Request().Context(), "Get")
	defer span.End()
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

func UpdateUser(c echo.Context) error {
	_, span := otel.Tracer(name).Start(c.Request().Context(), "Update")
	defer span.End()
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

func DeleteUser(c echo.Context) error {
	_, span := otel.Tracer(name).Start(c.Request().Context(), "Delete")
	defer span.End()
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}

func Healthz(c echo.Context) error {
	counter.Inc()
	_, span := otel.Tracer(name).Start(c.Request().Context(), "Health")
	defer span.End()
	return c.JSON(http.StatusOK, 200)
}
