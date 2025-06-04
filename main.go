package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Home page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		renderTemplate(w, "index.html")
	})

	// About page
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		renderPartial(w, "about.html")
	})

	// Projects page
	http.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		renderPartial(w, "projects.html")
	})

	// Experience page
	http.HandleFunc("/experience", func(w http.ResponseWriter, r *http.Request) {
		renderPartial(w, "experience.html")
	})

	// Leadership page
	http.HandleFunc("/leadership", func(w http.ResponseWriter, r *http.Request) {
		renderPartial(w, "leadership.html")
	})

	// Contact page
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		renderPartial(w, "contact.html")
	})

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func renderPartial(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}