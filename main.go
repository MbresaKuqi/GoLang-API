package main

import (
	"example.com/m/controller"
	"example.com/m/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()
	storage.NewDB()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("api/countries", controller.GetCountries)
	e.GET("/api/countries/:countryName", controller.CreateCountry)
	e.GET("/api/country/:countryName", controller.DeleteCountry)

	// Start server
	e.Logger.Fatal(e.Start(":3307"))
}
