package database

import (
	"database/sql"
	"log"
)

func Create(db *sql.DB) (err error) {
	userTable := `CREATE TABLE IF NOT EXISTS ip (country text, city text, lat float, lon float, isp text, query text);`
	statement, err := db.Prepare(userTable)

	if err != nil {
		log.Println(err.Error())

	}

	if _, err = statement.Exec(); err != nil {
		log.Println(err.Error())

	}

	return statement.Close()
}
