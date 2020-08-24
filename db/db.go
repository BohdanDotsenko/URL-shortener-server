package db

import (
	"database/sql"
	"fmt"
	"log"
)

// URL struct like in database
type URL struct {
	LongLink  string
	ShortLink string
}

//PrepeareDb prepeare database
func PrepeareDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "db/db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS URL (LongLink TEXT, ShortLink TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	return db, err
}

// NewLink adding
func NewLink(link URL, db *sql.DB) error {
	row, err := db.Query("SELECT * FROM URL WHERE LongLink=?", link.LongLink)
	if err != nil {

	}
	fmt.Println(row)
	return err
}
