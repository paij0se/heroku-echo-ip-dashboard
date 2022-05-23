package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/paij0se/ip/src/controllers"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Static("/", "src/public")
	e.POST("/u", controllers.UpdateData)
	e.GET("/ip", controllers.GetIp)
	e.GET("/ip/all", controllers.ReturnIps)

	port := os.Getenv("PORT")

	if port == "" {
		log.Println("The port to use is not declared, using port 8080.")

		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))

}
