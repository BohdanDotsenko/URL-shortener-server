package db

import (
	"database/sql"
	"fmt"
	"log"
)

// URL struct like in database
type URL struct {
	LongURL  string
	ShortURL string
}

//PrepeareDb prepeare database
func PrepeareDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "db/db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS URL (LongURL TEXT, ShortURL TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	return db, err
}

// NewURL adding
func NewURL(link string, database *sql.DB) error {
	row, err := database.Query("SELECT * FROM URL WHERE LongURL=?", link)
	if err != nil {
		statement, _ := database.Prepare("INSERT INTO URL (LongURL) VALUES (?)")
		statement.Exec()
		fmt.Printf("here\n")
	}
	defer row.Close()
	fmt.Printf("no here\n")
	return err
}
