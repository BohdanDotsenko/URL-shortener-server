package db

import (
	"database/sql"
	"log"
)

type links struct {
	LongLink  string
	ShotrLink string
}

//PrepeareDb prepeare database
func PrepeareDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "db/db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS links (LongLink TEXT, ShotrLink TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	return db, err
}
