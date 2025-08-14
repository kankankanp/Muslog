package presenters

import (
	"github.com/labstack/echo/v4"
)

func ErrorResponse(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, echo.Map{"error": message})
}
