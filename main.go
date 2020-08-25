package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BohdanDotsenko/URL-shortener-server/db"
	"github.com/BohdanDotsenko/URL-shortener-server/server"
	_ "github.com/mattn/go-sqlite3"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "URL:", r.URL.String())
}

func main() {
	database, err := db.PrepeareDb()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
	var url = db.URL{LongURL: "adwssgdasdgsgvx", ShortURL: "safasfaf"}

	server.OpenServer(database)
}
