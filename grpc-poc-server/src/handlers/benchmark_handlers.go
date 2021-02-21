package handlers

import (
	"io/ioutil"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func GetPayloadHandler(c echo.Context) error {
	size := c.Param("size")
	dat, err := ioutil.ReadFile("fixtures/" + size)
	check(err)
	return c.JSON(http.StatusOK, string(dat))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
