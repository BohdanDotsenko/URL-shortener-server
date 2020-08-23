package main

import (
	"fmt"

	"github.com/BohdanDotsenko/URL-shortener-server/tree/master/db"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Hello")
	db, err := db.OpenDb()
	defer db.Close()

	srv, err := newServer(db)

	// srv, err := newServer()
}
