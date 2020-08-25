package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"time"
	"math/rand"
	
	"github.com/BohdanDotsenko/URL-shortener-server/db"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)


func generateShorlLink() string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 5)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func HTMLHandler(w http.ResponseWriter, r *http.Request) {
	var Links db.Links
	err := r.ParseForm()
	if err != nil {
			log.Fatalln(err)
	}
	Links.LongURL = r.Form.Get("original")
	if Links.LongURL != "" && !db.ExistURL(Links.LongURL){
		fmt.Println(Links.LongURL)
		Links.ShortURL = generateShorlLink()
		db.NewURL(Links)
	}
	t, err := template.ParseFiles("./frontend/index.html")
		if err != nil {
			log.Fatalln(err)
		}

		t.Execute(w, Links)
}

func main() {
	err := db.PrepeareDb()
	if err != nil {
		log.Fatal(err)
	}
	router := mux.NewRouter()
    // router.HandleFunc("/products/{id:[0-9]+}", productsHandler)
    // router.HandleFunc("/articles/{id:[0-9]+}", productsHandler)
	router.HandleFunc("/", HTMLHandler).Methods("GET")
	router.HandleFunc("/", HTMLHandler).Methods("POST")
    http.Handle("/",router)
 
    fmt.Println("Server is listening port 8181")
    http.ListenAndServe(":8181", router)
}
