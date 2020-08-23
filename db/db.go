package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type table struct {
	LongLink  string
	ShotrLink string
}

func OpenDb() error {
	db, err := sql.Open("sqlite3", "db/db.sqlite3")
	if err != nil {
		panic(err)
	}
	err = AddURL(db, URL{Link: "https://pkg.go.dev/"})

	fmt.Println("db")
	return err
}
