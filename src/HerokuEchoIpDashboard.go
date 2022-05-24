package HerokuEchoIpDashboard

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/paij0se/heroku-echo-ip-dashboard/src/controllers"
	"github.com/paij0se/heroku-echo-ip-dashboard/src/download"
)

func HerokuEchoIpDashboard(e *echo.Echo) {
	download.Download()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // If you want restrict access to some domains, add them here
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Static("/dashboard", "herokudashboard/herokudashboard/public") // Also you can change the path of the static files
	e.POST("/u", controllers.UpdateData)
	e.GET("/ip", controllers.GetIp)
	e.GET("/ip/all", controllers.ReturnIps)

}
