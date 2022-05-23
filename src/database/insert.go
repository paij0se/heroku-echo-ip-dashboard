package database

import (
	"database/sql"
	"log"
)

func Insert(db *sql.DB, country string, city string, lat float64, lon float64, isp string, query string) (err error) {
	insert, err := db.Prepare("INSERT INTO ip (country, city, lat, lon, isp, query) VALUES ($1, $2, $3, $4, $5, $6)")

	if err != nil {
		log.Println(err.Error())

	}

	if _, err = insert.Exec(country, city, lat, lon, isp, query); err != nil {
		log.Println(err.Error())

	}

	return insert.Close()
}
