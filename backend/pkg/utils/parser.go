package utils

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func ParseID(c echo.Context, paramName string) (uint, error) {
	idStr := c.Param(paramName)
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
