package rest

import (
	echo "github.com/labstack/echo/v4"
	controllers "justino.com/poc/controllers"
)

func StartupRestServer() {
	e := echo.New()

	// Routes
	controllers.MapRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":10001"))
}
