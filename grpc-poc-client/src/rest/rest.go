package rest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetPayload(size string) {
	response, err := http.Get("http://localhost:10001/payload/" + size)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// responseData, err := ioutil.ReadAll(response.Body)
	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
}
