package main

import (
	"html/template"
	"log"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
    t, err := template.ParseFiles("templates/" + tmpl + ".html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    t.Execute(w, nil)
}

func main() {
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        renderTemplate(w, "index")
    })

    http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
        renderTemplate(w, "about")
    })

    http.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
        renderTemplate(w, "projects")
    })

    http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
        renderTemplate(w, "contact")
    })

    log.Println("Server running at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
