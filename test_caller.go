package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.Formatter = new(logrus.JSONFormatter)
	log.Formatter = new(logrus.TextFormatter) // default
	log.Level = logrus.DebugLevel
}

func main() {
	fmt.Println("Main function v2 :")
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Routes
	e.GET("/caller-test", goWithDocker)
	e.Logger.Fatal(e.Start(":8081"))
}

func goWithDocker(c echo.Context) error {
	fmt.Println("Handling call :")
	response, err := http.Get("http://test-testgo:8080/go-docker")
	//response, err := http.Get("http://localhost:8080/go-docker")
	if err != nil {
		fmt.Println("Error calling callee")
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		bodyString := string(bodyBytes)
		return c.JSON(http.StatusOK, bodyString)
	}

	return c.JSON(response.StatusCode, "Error occurred")
}
