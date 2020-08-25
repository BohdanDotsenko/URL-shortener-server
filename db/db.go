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
	prepeare := `CREATE TABLE IF NOT EXISTS URL (LongURL TEXT, ShortURL TEXT)`
	statement, err := db.Prepare(prepeare)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	return db, err
}

// ExistURL ?
func ExistURL(str string, database *sql.DB) bool {
	exist := `SELECT LongURL FROM URL WHERE LongURL = ?`
	err := database.QueryRow(exist, str).Scan(&str)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return false
	}
	return true
}

// NewURL adding
func NewURL(URL URL, database *sql.DB) error {
	if ExistURL(URL.LongURL, database) {
		fmt.Printf("URL already exist in database\n")
		return nil
	}
	insertSQL := `INSERT INTO URL (LongURL, ShortURL) VALUES (?, ?)`
	statement, err := database.Prepare(insertSQL)
	statement.Exec(URL.LongURL, URL.ShortURL)
	fmt.Printf("URL successfully added in database\n")
	return err
}

func generateID() string {
	b := make([]byte, 5)
	// for i := range b {
	// 	b[i] = symbols[rand.Int63()%int64(len(symbols))]
	// }
	return string(b)
}

//GetShortURL from database
func GetShortURL(longURL string) (string, error) {

	return "", nil
}
