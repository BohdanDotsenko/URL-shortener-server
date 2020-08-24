package main

import (
	"fmt"
	"log"

	"github.com/BohdanDotsenko/URL-shortener-server/db"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Hello")
	db, err := db.PrepeareDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	NewURL("URL", db)
	// srv, err := newServer(db)

	// srv, err := newServer()
}
