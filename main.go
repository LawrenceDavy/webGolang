package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

var tpl *template.Template

type pageData struct {
	Title string
	Date string
	Score int
	Descripton string
}



func init() {
	tpl = template.Must(template.ParseGlob("template/*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	pd := pageData {
		Title: "Index Page",
	}
	err := tpl.ExecuteTemplate(w, "index.html", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	pd := pageData {
		Title: "About Page",
	}
	err := tpl.ExecuteTemplate(w, "about.html", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
