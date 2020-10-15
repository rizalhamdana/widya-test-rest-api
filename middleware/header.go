package middleware

import (
	"github.com/labstack/echo/v4"
)

// ServerHeader function
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "master/1.0 (AArch64)")
		return next(c)
	}
}
