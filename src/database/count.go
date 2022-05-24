package database

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func Rows(db *sql.DB, c echo.Context) (err error) {
	rows, err := db.Query("SELECT country FROM ip")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var country string
		err = rows.Scan(&country)
		if err != nil {
			log.Println(err.Error())
		}
		if err != nil {
			log.Println(err.Error())
		}
		json.NewEncoder(c.Response()).Encode(country)
	}
	return rows.Close()
}
