package handlers

import (
	"log"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func GetLargePayloadHandler(c echo.Context) error {
	log.Printf("Payload grandão")
	return c.JSON(http.StatusOK, "Payload grandão")
}
