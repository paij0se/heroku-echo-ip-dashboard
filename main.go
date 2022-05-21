package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Ip struct {
	Country string
	City    string
	Lat     float64
	Lon     float64
	Isp     string
	Query   string
}

func main() {
	e := echo.New()
	e.GET("/ip", func(c echo.Context) error {
		ip := c.RealIP()
		req, err := http.Get("http://ip-api.com/json/" + ip)
		if err != nil {
			log.Fatal(err)
		}
		body, err := ioutil.ReadAll(req.Body)
		var ipUser Ip
		json.Unmarshal(body, &ipUser)
		return json.NewEncoder(c.Response()).Encode(map[string]string{"country": ipUser.Country, "query": ipUser.Query, "isp": ipUser.Isp, "city": ipUser.City})
	})
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Static("/", "public")
	e.Logger.Fatal(e.Start(":3456"))

}
