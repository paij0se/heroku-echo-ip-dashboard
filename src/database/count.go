package database

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func Rows(db *sql.DB, c echo.Context) {
	rows, err := db.Query("SELECT * FROM ip")
	if err != nil {
		log.Println(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var country, city, isp, query string
		var lat, lon float64
		err = rows.Scan(&country, &city, &lat, &lon, &isp, &query)
		if err != nil {
			log.Println(err.Error())
		}
		if err != nil {
			log.Println(err.Error())
		}
		c.JSON(200, map[string]interface{}{
			"country": country,
			"city":    city,
			"ip":      query,
		})

	}
}
