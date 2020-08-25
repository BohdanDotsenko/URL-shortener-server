package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	_ "github.com/mattn/go-sqlite3"
	"github.com/BohdanDotsenko/URL-shortener-server/db"
	"html/template"
	"log"
)

type TestCase struct {
	Link string
	StatusCode int
}

func TestRedirectHandler(t *testing.T) {
	cases := []TestCase{
		// TestCase{
		// 	Link: "1234567",
		// 	StatusCode: http.StatusOK,
		// },
		TestCase{
			Link: "/",
			StatusCode: http.StatusNotFound,
		},
	}
	for caseNum, item := range cases {
		// cases[caseNum].Link = "/"
		req := httptest.NewRequest("GET", cases[caseNum].Link, nil)
		w := httptest.NewRecorder()
		RedirectHandler(w, req)
		if w.Code != item.StatusCode {
			t.Errorf("[%d] wrong StatusCode: got %d, expected %d",
				caseNum, w.Code, item.StatusCode)
		}
	}
	fmt.Println("PASS----TestRedirectHandler")
}



func TestHTMLHandlerTest(t *testing.T) {
	cases := []TestCase{
		TestCase{
			Link: "https://www.google.com.ua/",
			StatusCode: http.StatusOK,
		},
	}
	for caseNum, item := range cases {
		// cases[caseNum].Link = "/"
		req := httptest.NewRequest("POST", cases[caseNum].Link, nil)
		w := httptest.NewRecorder()
		HTMLHandlerTest(w, req)
		if w.Code != item.StatusCode {
			t.Errorf("[%d] wrong StatusCode: got %d, expected %d",
				caseNum, w.Code, item.StatusCode)
		}
		fmt.Println("PASS----TestHTMLHandler")
	}
}

func HTMLHandlerTest(w http.ResponseWriter, r *http.Request) {
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
	t, err := template.ParseFiles("./../frontend/index.html")
	if err != nil {
		log.Fatalln(err)
	}

	t.Execute(w, Links)
}