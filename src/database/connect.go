package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Connect() (db *sql.DB, err error) {
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Println(err.Error())

	}

	if err = Create(db); err != nil {
		log.Println(err.Error())

	}

	return
}
