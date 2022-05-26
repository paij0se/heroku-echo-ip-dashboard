package HerokuEchoIpDashboard

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/paij0se/heroku-echo-ip-dashboard/src/controllers"
	"github.com/paij0se/heroku-echo-ip-dashboard/src/download"
	"golang.org/x/time/rate"
)

// This Is the main function that will be called by Heroku, and with a rete-Limiter
func HerokuEchoIpDashboardWithRateLimiter(e *echo.Echo, RateLimiterMemoryStore float64, rateN int, burst int, ExpireTimeSec int64) {
	download.Download()
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(RateLimiterMemoryStore))))
	e.Use(middleware.Recover())
	config := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: rate.Limit(rateN), Burst: burst, ExpiresIn: time.Duration(ExpireTimeSec) * time.Second},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, nil)
		},
	}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // If you want restrict access to some domains, add them here
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Static("/dashboard", "herokudashboard/herokudashboard/public") // Also you can change the path of the static files
	e.POST("/ip/update", controllers.UpdateData, middleware.RateLimiterWithConfig(config))
	e.GET("/ip", controllers.GetIp, middleware.RateLimiterWithConfig(config))
	e.GET("/ip/all", controllers.ReturnIps, middleware.RateLimiterWithConfig(config))

}

func HerokuEchoIpDashboard(e *echo.Echo) {
	download.Download()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // If you want restrict access to some domains, add them here
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Static("/dashboard", "herokudashboard/herokudashboard/public") // Also you can change the path of the static files
	e.POST("/ip/update", controllers.UpdateData)
	e.GET("/ip", controllers.GetIp)
	e.GET("/ip/all", controllers.ReturnIps)
}
