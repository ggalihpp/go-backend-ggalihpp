package main

import (
	"net/http"

	"github.com/ggalihpp/go-backend-ggalihpp/minio"
	example "github.com/ggalihpp/go-backend-ggalihpp/route-example"
	"github.com/labstack/echo"
)

func setupHandlers(e *echo.Echo) {
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	exampleRoute := e.Group("/example")
	example.SetupHandler(exampleRoute)

	minioRoute := e.Group("/file")
	minio.SetupHandler(minioRoute)
}
