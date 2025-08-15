package server

import (
	"github.com/labstack/echo/v4"
	"log"
)

func StartServer(e *echo.Echo) {
	log.Fatal(e.Start(":8080"))
}