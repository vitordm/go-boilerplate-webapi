package main

import (
	"github.com/labstack/echo/v4"
	"github.concur.com/I573758/example-golang-webapi/internal/app"
)

func main() {
	echoRouter := echo.New()
	app.Configure(echoRouter)
	app.Start(echoRouter)
}
