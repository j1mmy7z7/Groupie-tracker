package main

import (
	"groupie-tracker/handlers"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handlers.Homehandler)
	http.HandleFunc("/display", handlers.Locationhandler)

	log.Println("running on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
