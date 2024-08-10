package handlers

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"

	"groupie-tracker/data"
)

var (
	tpl *template.Template
	err error
)

func init() {
	tpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		return
	}
}

func Rendertemplate(w http.ResponseWriter, data interface{}) {
	err = tpl.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Homehandler(w http.ResponseWriter, r *http.Request) {
	url := "https://groupietrackers.herokuapp.com/api/artists"

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "error fetching data", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()
	// Check for successful response
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Unexpected status code from external service", http.StatusInternalServerError)
		return
	}

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading reponse data", http.StatusInternalServerError)
		return
	}
	var Bandis []data.Band

	json.Unmarshal(resBody, &Bandis)

	data := data.PageData{
		Title: "Home",
		Bands: Bandis,
	}
	Rendertemplate(w, data)
}
