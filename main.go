package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"ibp/interfaces"
)

func main () {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	interfaces.Routes(e)
	interfaces.Run(e, 8000)
}