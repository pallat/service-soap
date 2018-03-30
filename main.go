//   //go:generate go build -o mgo2es main.go

package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/service", Service)
	e.GET("/service", ServiceWSDL)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
