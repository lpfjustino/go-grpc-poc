package controllers

import (
	echo "github.com/labstack/echo/v4"
	handlers "justino.com/poc/handlers"
)

func MapRoutes(e *echo.Echo) {
	e.GET("/large-payload", handlers.GetLargePayloadHandler)
}
