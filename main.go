package main

import (
	"fmt"

	"./db"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Hello")
	db, err := db.OpenDb()
	defer db.Close()

	srv, err := newServer(db)

	// srv, err := newServer()
}
