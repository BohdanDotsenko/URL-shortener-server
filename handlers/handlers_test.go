package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	_ "github.com/mattn/go-sqlite3"
	"github.com/BohdanDotsenko/URL-shortener-server/db"
)

type TestCase struct {
	Link string
	StatusCode int
}

func TestRedirectHandler(t *testing.T) {
	var Links db.Links
	Links.ShortURL = "1234567"
	Links.LongURL = "https://www.google.com.ua/"
	db.NewURL(Links)
	cases := []TestCase{
		TestCase{
			Link: "1234567",
			StatusCode: http.StatusOK,
		},
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
}
