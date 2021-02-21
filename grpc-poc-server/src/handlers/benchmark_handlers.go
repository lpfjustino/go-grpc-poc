package handlers

import (
	"io/ioutil"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func GetLargePayloadHandler(c echo.Context) error {
	dat, err := ioutil.ReadFile("fixtures/1mb")
	check(err)
	return c.JSON(http.StatusOK, string(dat))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
