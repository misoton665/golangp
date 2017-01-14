package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/echo", func(c echo.Context) error {
		return c.String(http.StatusOK, "echo")
	})
	e.Logger.Fatal(e.StartAutoTLS(":1323"))
}
