package handlers

import (
	"encoding/json"
	"groupie-tracker/data"
	"io"
	"net/http"
)

// Locationhandler processes form submissions and fetches data from a URL.
func Locationhandler(w http.ResponseWriter, r *http.Request) {
	// Parse form values
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	formType := r.FormValue("form_type")
	url := r.FormValue("url")

	// Fetch data from the URL
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Initialize data structures
	var location data.Location
	var dates data.Date
	var relations data.Relation

	// Check for successful response
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Unexpected status code from external service", http.StatusInternalServerError)
		return
	}

	// Read and parse the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		return
	}

	switch formType {
	case "Dates":
		if err := json.Unmarshal(body, &dates); err != nil {
			http.Error(w, "Error unmarshalling dates data", http.StatusInternalServerError)
			return
		}
		data := data.PageData{
			Title: "Dates",
			Bands: dates,
		}
		Rendertemplate(w, data)

	case "Location":
		if err := json.Unmarshal(body, &location); err != nil {
			http.Error(w, "Error unmarshalling location data", http.StatusInternalServerError)
			return
		}
		data := data.PageData{
			Title: "Location",
			Bands: location,
		}
		Rendertemplate(w, data)

	case "Relations":
		if err := json.Unmarshal(body, &relations); err != nil {
			http.Error(w, "Error unmarshalling relations data", http.StatusInternalServerError)
			return
		}
		data := data.PageData{
			Title: "Relations",
			Bands: relations,
		}
		Rendertemplate(w, data)

	default:
		http.Error(w, "Unknown form type", http.StatusBadRequest)
	}
}
