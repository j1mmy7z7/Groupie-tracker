package main

import (
	"groupie-tracker/data"
	"groupie-tracker/handlers"
	"log"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		switch path {
		case "/":
			handlers.Homehandler(w, r)
		case "/display":
			handlers.Locationhandler(w, r)
		default:
			data := data.PageData{
				Title: "Error",
				Bands: struct {
					Message string
					Code    int
				}{
					Message: "Page Not Found",
					Code:    404,
				},
			}
			handlers.Rendertemplate(w, data)
		}
	})

	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err.Error())
		}
	}()

	log.Println("running on 8080")

	select {}

}
