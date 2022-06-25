package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	//Instance Echo
	e := echo.New()

	//routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")

	})
	e.Logger.Fatal(e.Start(":1323"))
}
