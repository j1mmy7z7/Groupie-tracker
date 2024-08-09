package main

import(
	"net/http"
	"log"
	"html/template"
	"groupie-tracker/handlers"
)

var tpl *template.Template


func main(){

	//tpl, _ := template.ParseGlob("templates/*.html")
	http.HandleFunc("/", handlers.Homehandler)

	log.Println("running on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}