package main

import (
	"fmt"
	"log"

	"github.com/BohdanDotsenko/URL-shortener-server/db"
	_ "github.com/mattn/go-sqlite3"
)

type URL struct {
	LongLink  string
	ShortLink string
}

func main() {
	fmt.Println("Hello")
	db, err := db.PrepeareDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var URL URL = URL{"dasdasd", "sadasdas"}
	NewLink(URL, db)
	// srv, err := newServer(db)

	// srv, err := newServer()
}
