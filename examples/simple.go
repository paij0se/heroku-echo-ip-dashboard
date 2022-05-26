package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	HerokuEchoIpDashboard "github.com/paij0se/heroku-echo-ip-dashboard/src"
	re "github.com/paij0se/heroku-echo-ip-dashboard/src/controllers"
)

func main() {
	e := echo.New()
	HerokuEchoIpDashboard.HerokuEchoIpDashboardWithRateLimiter(e, 10, 10, 2, 120)
	//HerokuEchoIpDashboard.HerokuEchoIpDashboard(e)

	e.GET("/", func(c echo.Context) error {
		re.Requester(c.Scheme() + "://" + c.Request().Host)
		return c.File("examples/index.html")
	})

	port := os.Getenv("PORT")

	if port == "" {
		log.Println("The port to use is not declared, using port 8080")

		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))

}
