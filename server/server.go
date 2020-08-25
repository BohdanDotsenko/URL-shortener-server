package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/BohdanDotsenko/URL-shortener-server/db"
	"github.com/julienschmidt/httprouter"
)

type handle struct {
}

// OpenServer starting and working server
func OpenServer(database *sql.DB) {
	server := http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("starting server at :8080")
	server.ListenAndServe()
	router := httprouter.New()
	router.GET("/:id", handleForm())

	fmt.Print("sadas")
}

func handleForm() httprouter.Handle {
	var URL db.URL

}
