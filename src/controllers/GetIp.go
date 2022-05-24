package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/paij0se/heroku-echo-ip-dashboard/src/database"
	"github.com/paij0se/heroku-echo-ip-dashboard/src/interfaces"
)

func Post(c echo.Context, Country string, City string, Lat float64, Lon float64, Isp string, Query string) error {
	postBody, _ := json.Marshal(map[string]string{
		"Country": Country,
		"City":    City,
		"Lat":     fmt.Sprintf("%f", Lat),
		"Lon":     fmt.Sprintf("%f", Lon),
		"Isp":     Isp,
		"Query":   Query,
	})
	fmt.Println(string(postBody))
	resp, err := http.Post(c.Scheme()+"://"+c.Request().Host+"/u", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
	return nil

}

func ReturnIps(c echo.Context) error {
	db, _ := database.Connect()
	database.Rows(db, c)
	return nil
}

/*
This little function is going to make a get request to your-app.com/ip
and return the ip of the user And insert it in the database
*/
func Requester(url string) {
	req, err := http.Get(url + "/ip")
	if err != nil {
		log.Fatal(err)
	}
	body, _ := ioutil.ReadAll(req.Body)
	log.Println(string(body))
}

func GetIp(c echo.Context) error {
	ip := c.RealIP()
	req, err := http.Get("http://ip-api.com/json/" + ip)
	if err != nil {
		log.Fatal(err)
	}
	body, _ := ioutil.ReadAll(req.Body)
	var ipUser interfaces.Ip
	json.Unmarshal(body, &ipUser)
	Post(c, ipUser.Country, ipUser.City, ipUser.Lat, ipUser.Lon, ipUser.Isp, ip)
	return json.NewEncoder(c.Response()).Encode(map[string]string{"country": ipUser.Country, "query": ipUser.Query, "isp": ipUser.Isp, "city": ipUser.City})
}
