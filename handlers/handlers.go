package handlers

import (
	"fmt"
	"github.com/BohdanDotsenko/URL-shortener-server/db"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)


func HTMLHandler(w http.ResponseWriter, r *http.Request) {
	var Links db.Links
	err := r.ParseForm()
	if err != nil {
		fmt.Println("handlers.go -> HTMLHandler -> ParseForm() error: ", err)
	}
	Links.LongURL = r.Form.Get("original")
	if Links.LongURL != "" {
		Links.ShortURL = generateShortLink()
		db.NewURL(Links)
	}
	t, err := template.ParseFiles("./frontend/index.html")
	if err != nil {
		fmt.Println("handlers.go -> HTMLHandler -> ParseFiles() error: ", err)

	}

	t.Execute(w, Links)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	var Links db.Links
	vars := mux.Vars(r)
	Links.ShortURL = vars["id"]
	Links.LongURL = db.GetLongURL(Links.ShortURL)
	if Links.LongURL == "" {
		http.NotFound(w, r)
	} else {
		http.Redirect(w, r, Links.LongURL, http.StatusFound)
	}
}
