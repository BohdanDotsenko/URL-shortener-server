package db

import (
	"database/sql"
	"fmt"
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
		fmt.Println("db.go -> PrepareDB() -> sql.Open() error : ", err)
	}
	defer database.Close()
	prepeare := `CREATE TABLE IF NOT EXISTS Links (LongURL TEXT, ShortURL TEXT)`
	statement, err := database.Prepare(prepeare)
	if err != nil {
		fmt.Println("db.go -> PrepeareDb() -> Prepare() error : ", err)
	}
	statement.Exec()
	return  err
}

// ExistURL ?
func ExistURL(str string) bool {
	database, err := sql.Open("sqlite3", "db/db.sqlite3")
	if err != nil {
		fmt.Println("db.go -> ExistURL() -> PrepareDB() error : ", err)
	}
	defer database.Close()
	exist := `SELECT LongURL FROM Links WHERE LongURL = ?`
	err = database.QueryRow(exist, str).Scan(&str)
	if err != nil {
		if err != sql.ErrNoRows {
			fmt.Println("db.go -> ExistURL() -> QueryRow() error: ", err)
		}
		return false
	}
	return true
}

// NewURL adding
func NewURL(Links Links) error {
	database, err := sql.Open("sqlite3", "db/db.sqlite3")
	if err != nil {
		fmt.Println("db.go -> NewURL() -> sql.Open() error: ", err)
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
		fmt.Println("db.go -> GetLongURL() -> sql.Open() error: ", err)
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
