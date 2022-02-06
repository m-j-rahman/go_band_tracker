package tracker

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// Handles main page with names and pictures of all artists.
func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.Method == "GET" {
		Err("404 Not Found", http.StatusNotFound, w)
		return
	}

	if r.URL.Path == "/" && r.Method != "GET" {
		Err("405 Method Not Allowed", http.StatusMethodNotAllowed, w)
		return
	}

	htmlTemplate, err := template.ParseFiles("templates/index.html")
	if err != nil {
		Err("500 Internal Server Error", http.StatusInternalServerError, w)
		return
	}

	GettingAPIData(w)

	htmlTemplate.Execute(w, ArtistsInfo)
}

// Handles pages with detailed info about specific artists.
func ArtistPage(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path) < 10 || (r.URL.Path[:9] != "/artists/") {
		Err("404 Not Found", http.StatusNotFound, w)
		return
	}

	if r.Method != "GET" {
		Err("405 Method Not Allowed", http.StatusMethodNotAllowed, w)
		return
	}

	htmlTemplate, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		Err("500 Internal Server Error", http.StatusInternalServerError, w)
		return
	}

	ArtistID := strings.TrimPrefix(r.URL.Path, "/artists/")

	ID, err := strconv.Atoi(ArtistID)
	if err != nil {
		Err("400 Bad Request", http.StatusBadRequest, w)
		return
	}

	GettingAPIData(w)

	if len(ArtistsInfo) < ID {
		Err("404 Not Found", http.StatusNotFound, w)
		return
	} else if ID < 1 {
		Err("400 Bad Request", http.StatusBadRequest, w)
		return
	}

	err = htmlTemplate.Execute(w, ArtistsInfo[ID-1])

	if err != nil {
		Err("500 Internal Server Error", http.StatusInternalServerError, w)
		return
	}
}

// Calls handle functions.
func MainHandler() {
	http.HandleFunc("/", MainPage)
	http.HandleFunc("/artists/", ArtistPage)
	http.ListenAndServe(":8080", nil)
}
