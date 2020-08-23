package db

import (
	"database/sql"
	"fmt"
)

type table struct {
	LongLink  string
	ShotrLink string
}

//OpenDb create or open database
func OpenDb() error {
	db, err := sql.Open("sqlite3", "db/db.sqlite3")
	if err != nil {
		panic(err)
	}
	fmt.Println("db")
	return err
}
