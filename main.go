package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type pageData struct {
	Title string
	FirstName string
}

func init() {
	tpl = template.Must(template.ParseGlob("template/*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/apply", apply)
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

func contact(w http.ResponseWriter, r *http.Request) {
	pd := pageData {
		Title: "Contact Page",
	}
	err := tpl.ExecuteTemplate(w, "contact.html", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func apply(w http.ResponseWriter, r *http.Request) {

	pd := pageData {
		Title: "Apply Page",
	}

	var first string
	if r.Method == http.MethodPost {
		first = r.FormValue("fname")
		pd.FirstName = first
	}
	err := tpl.ExecuteTemplate(w, "apply.html", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
