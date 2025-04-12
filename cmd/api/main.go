package main

import (
	"github.com/8bury/sumup4me/internal/config"
	"github.com/labstack/echo/v4"
)

func main() {
	api := echo.New()
	
	config.ConfigureApi(api)

	api.Start(":8080")
}