package main

import (
	"context"
	"log"
	"os"

	"github.com/eugenekatsov/go-sample-service/internal"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
)

func main() {

	e := echo.New()

	l := log.New(os.Stdout, "", 0)

	// Write telemetry data to a file.
	f, err := os.Create("traces.txt")
	if err != nil {
		l.Fatal(err)
	}
	defer f.Close()

	exp, err := internal.NewExporter(f)
	if err != nil {
		l.Fatal(err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(internal.NewResource()),
	)
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			l.Fatal(err)
		}
	}()

	otel.SetTracerProvider(tp)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(echoprometheus.NewMiddleware("gss"))

	// Routes
	e.POST("/users", internal.CreateUser)
	e.GET("/healthz", internal.Healthz)
	e.GET("/users/:id", internal.GetUser)
	e.GET("/metrics", echoprometheus.NewHandler())
	e.PUT("/users/:id", internal.UpdateUser)
	e.DELETE("/users/:id", internal.DeleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":80"))
}
