package database

import (
	"database/sql"
	"log"

	"github.com/paij0se/ip/src/interfaces"
)

func Query(db *sql.DB, API *interfaces.Ip) (err error) {
	row, err := db.Query("SELECT country, city, lat, lon, isp, query FROM ip WHERE query = $1", API.Query)

	if err != nil {
		log.Println(err.Error())

		return
	}

	if row.Next() {
		err = row.Scan(&API.Country, &API.City, &API.Lat, &API.Lon, &API.Isp, &API.Query)

		if err != nil {
			log.Println(err.Error())

			return
		}

	}

	return row.Close()
}
