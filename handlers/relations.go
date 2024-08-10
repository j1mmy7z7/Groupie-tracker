package handlers

import (
	"encoding/json"
	"groupie-tracker/data"
	"io"
	"net/http"
)

func Relationshandler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Error fetching dates data", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		Body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Error reading dated data", http.StatusInternalServerError)
			return
		}
		var relation data.Relation
		json.Unmarshal(Body, &relation)
		data := data.PageData{
			Title: "Relations",
			Bands: relation,
		}
		Rendertemplate(w, data)
	}

}
