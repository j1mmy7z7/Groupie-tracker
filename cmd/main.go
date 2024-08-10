package main

import (
	"groupie-tracker/handlers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handlers.Homehandler)
	http.HandleFunc("/location", handlers.Locationhandler)
	http.HandleFunc("/dates", handlers.Dateshandler)
	http.HandleFunc("/relation", handlers.Relationshandler)

	log.Println("running on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
