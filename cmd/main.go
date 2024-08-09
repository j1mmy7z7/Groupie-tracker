package main

import(
	"net/http"
	"log"
	"html/template"
)

var tpl *template.Template

func main(){

	tpl, _ := template.ParseGlob("templates/*.html")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		tpl.ExecuteTemplate(w, "base.html", nil)
	})

	log.Println("running on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}