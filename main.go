package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BohdanDotsenko/URL-shortener-server/db"
	"github.com/BohdanDotsenko/URL-shortener-server/handlers"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := db.PrepeareDb()
	if err != nil {
		log.Fatal(err)
	}
	router := mux.NewRouter()
	router.HandleFunc("/{id:[0-9a-zA-Z]+}", handlers.RedirectHandler).Methods("GET")
	router.HandleFunc("/", handlers.HTMLHandler).Methods("GET")
	router.HandleFunc("/", handlers.HTMLHandler).Methods("POST")
	http.Handle("/", router)

	fmt.Println("Server is listening port 8181")
	log.Fatal(http.ListenAndServe(":8181", router))
}
