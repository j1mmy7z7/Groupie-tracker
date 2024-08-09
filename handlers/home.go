package handlers

import (
	//"encoding/json"
	"encoding/json"
	//"fmt"
	"html/template"
	"io"
	"log"

	//"log"
	"groupie-tracker/data"
	"net/http"
)

var tpl *template.Template
var err error

func init() {
	tpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		return
	}
} 

func Rendertemplate(w http.ResponseWriter, data interface{}) error {
	err = tpl.ExecuteTemplate(w, "base.html", data)

	return err
}

func Homehandler(w http.ResponseWriter, r *http.Request) {
	url := "https://groupietrackers.herokuapp.com/api/artists"

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "error fetching data", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		resBody , err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Error reading reponse data", http.StatusInternalServerError)
			return
		}

		var Bandis []data.Band

		json.Unmarshal(resBody, &Bandis)

		data := data.PageData{
			Title: "Header Set",
			Bands: Bandis,
		}
		err = Rendertemplate(w, data)

		if err != nil {
			log.Println(err.Error())
			return
		}

	}
}
