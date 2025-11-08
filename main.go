package main

import (
	"github.com/labstack/echo/v4"
	"github.com/vitordm/go-boilerplate-webapi/internal/app"
)

func main() {
	echoRouter := echo.New()
	app.Configure(echoRouter)
	app.Start(echoRouter)
}
