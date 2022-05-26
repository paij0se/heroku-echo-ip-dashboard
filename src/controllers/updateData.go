package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/paij0se/heroku-echo-ip-dashboard/src/database"
	"github.com/paij0se/heroku-echo-ip-dashboard/src/interfaces"
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
	if ipUser.Country == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Country is required"})
	}
	if err = database.Insert(db, ipUser.Country, ipUser.City, ipUser.Lat, ipUser.Lon, ipUser.Isp, ipUser.Query); err != nil {
		log.Println(err.Error())
		return json.NewEncoder(c.Response()).Encode(map[string]string{"error": "Error inserting data"})
	}
	if err = db.Close(); err != nil {
		log.Println(err.Error())

	}
	if db.Close(); err != nil {
		log.Println(err.Error())

	}
	return json.NewEncoder(c.Response()).Encode(map[string]string{"success": "200 - OK"})
}
