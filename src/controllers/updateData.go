package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/paij0se/ip/src/database"
	"github.com/paij0se/ip/src/interfaces"
)

func UpdateData(c echo.Context) error {
	var ipUser interfaces.Ip
	reqBody, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Println(err.Error())
		return json.NewEncoder(c.Response()).Encode(map[string]string{"error": "Error reading request body"})
	}
	json.Unmarshal(reqBody, &ipUser)
	db, err := database.Connect()
	if err != nil {
		log.Println(err.Error())
		return json.NewEncoder(c.Response()).Encode(map[string]string{"error": "Error connecting to database"})
	}
	if err = database.Insert(db, ipUser.Country, ipUser.City, ipUser.Lat, ipUser.Lon, ipUser.Isp, ipUser.Query); err != nil {
		log.Println(err.Error())
		return json.NewEncoder(c.Response()).Encode(map[string]string{"error": "Error inserting data"})
	}
	/*
		{
			"Country": "United States",
			"City": "New York",
			"Lat": 40.7128,
			"Lon": -74.0060,
			"Isp": "Verizon",
			"Query": "182.255.1.2
		}
	*/
	if err = db.Close(); err != nil {
		log.Println(err.Error())

	}
	return json.NewEncoder(c.Response()).Encode(map[string]string{"success": "200 - OK"})
}
