package handlers

import (
	"encoding/json"
	"groupie-tracker/data"
	"io"
	"net/http"
)

func Locationhandler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Error fetching location data", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		Body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Error reading locations data", http.StatusInternalServerError)
			return
		}
		var location data.Location
		json.Unmarshal(Body, &location)
		data := data.PageData{
			Title: "location",
			Bands: location,
		}
		Rendertemplate(w, data)
	}

}
