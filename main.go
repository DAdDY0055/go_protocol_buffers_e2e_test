package main

import (
	"net/http"

	"github.com/DADDY0055/go_protocol_buffers_e2e_test/proto"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
