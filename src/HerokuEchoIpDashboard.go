package HerokuEchoIpDashboard

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/paij0se/heroku-echo-ip-dashboard/src/controllers"
)

/*
This little function is going to make a get request to your-app.com/ip
and return the ip of the user And insert it in the database
*/
func NumberOfTimesToCount(c echo.Context) error {
	req, err := http.Get(c.Scheme() + "://" + c.Request().Host + "/ip")
	if err != nil {
		log.Fatal(err)
	}
	body, _ := ioutil.ReadAll(req.Body)
	log.Println(string(body))
	return nil
}

func HerokuEchoIpDashboard(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // If you want restrict access to some domains, add them here
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Static("/dashboard", "src/public") // Also you can change the path of the static files
	e.POST("/u", controllers.UpdateData)
	e.GET("/ip", controllers.GetIp)
	e.GET("/ip/all", controllers.ReturnIps)

}
