package main

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	HerokuEchoIpDashboard "github.com/paij0se/heroku-echo-ip-dashboard/src"
)

func Test(testing *testing.T) {
	e := echo.New()
	HerokuEchoIpDashboard.HerokuEchoIpDashboard(e)

	e.GET("/", func(c echo.Context) error {
		go HerokuEchoIpDashboard.NumberOfTimesToCount(c)
		return c.String(http.StatusOK, "Hello, World!")
	})
	port := os.Getenv("PORT")

	if port == "" {
		log.Println("The port to use is not declared, using port 8080")

		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))

}
