package handlers

import (
	"fmt"
	"github.com/BohdanDotsenko/URL-shortener-server/db"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func generateShortLink() string {
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
	if Links.LongURL != "" {
		Links.ShortURL = generateShortLink()
		db.NewURL(Links)
	}
	t, err := template.ParseFiles("./frontend/index.html")
	if err != nil {
		log.Fatalln(err)
	}

	t.Execute(w, Links)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	var Links db.Links
	vars := mux.Vars(r)
	Links.ShortURL = vars["id"]
	Links.LongURL = db.GetLongURL(Links.ShortURL)
	fmt.Println("Links.LongURL : ")
	fmt.Println(Links.LongURL)
	if Links.LongURL == "" {
		http.NotFound(w, r)
	} else {
		http.Redirect(w, r, Links.LongURL, http.StatusFound)
	}
}
