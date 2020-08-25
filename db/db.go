package db

import (
	"database/sql"
	"fmt"
	"log"
)

// URL struct like in database
type Links struct {
	LongURL  string
	ShortURL string
}

//PrepeareDb prepeare database
func PrepeareDb() error {
	database, err := sql.Open("sqlite3", "db/db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	prepeare := `CREATE TABLE IF NOT EXISTS Links (LongURL TEXT, ShortURL TEXT)`
	statement, err := database.Prepare(prepeare)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	var Links Links
	Links.ShortURL = `123sa`
	Links.LongURL = `https://www.google.com.ua/`
	if ExistURL(Links.LongURL) {
		return err
	}
	NewURL(Links)
	return  err
}

// ExistURL ?
func ExistURL(str string) bool {
	database, err := sql.Open("sqlite3", "db/db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	exist := `SELECT LongURL FROM Links WHERE LongURL = ?`
	err = database.QueryRow(exist, str).Scan(&str)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return false
	}
	return true
}

// NewURL adding
func NewURL(Links Links) error {
	database, err := sql.Open("sqlite3", "db/db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	insertSQL := `INSERT INTO Links (LongURL, ShortURL) VALUES (?, ?)`
	statement, err := database.Prepare(insertSQL)
	statement.Exec(Links.LongURL, Links.ShortURL)
	fmt.Printf("URL successfully added in database\n")
	return err
}


//GetShortURL from database
func GetLongURL(shortURL string) string {
	var LongURL string
	database, err := sql.Open("sqlite3", "db/db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	row, err := database.Query("SELECT * FROM Links WHERE shortURL=?", shortURL)
	if err != nil {
		return ""
	}
	defer row.Close()
	for row.Next() {
		row.Scan(&LongURL, &shortURL)
	}
	return LongURL
}
