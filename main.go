package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", home)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal("listen and serve: ", err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "home.gohtml", nil)
}
