package main

import (
	"log"

	"github.com/BohdanDotsenko/URL-shortener-server/db"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database, err := db.PrepeareDb()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	var url = db.URL{LongURL: "asgfagagadgsadas", ShortURL: "safasfaf"}
	db.NewURL(url, database)
	// srv, err := newServer(db)

	// srv, err := newServer()
}
