package rest

import (
	"log"
	"net/http"

	echo "github.com/labstack/echo/v4"
	controllers "justino.com/poc/controllers"
)

func Teste(c echo.Context) error {
	log.Printf("Teste")
	return c.JSON(http.StatusOK, "Teste")
}

func StartupRestServer() {
	e := echo.New()

	// Routes
	controllers.MapRoutes(e)
	e.GET("/a", Teste)

	// Start server
	e.Logger.Fatal(e.Start(":10001"))
}
