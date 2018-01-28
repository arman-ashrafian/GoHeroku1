package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000" // gin port
	}

	// Mux Router
	r := mux.NewRouter()

	// serve static files
	r.PathPrefix("/static/").
		Handler(http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	// routes
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/signin", signinHandler)
	r.HandleFunc("/blog/{blogid}", blogHandler)

	log.Fatal(http.ListenAndServe(":"+port, r))

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(
		"templates/base.html",
		"templates/home.html",
	)
	t.ExecuteTemplate(w, "base", "")
}

func signinHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(
		"templates/base.html",
		"templates/sign_in.html",
	)
	t.ExecuteTemplate(w, "base", "")
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	t, _ := template.ParseFiles(
		"templates/base.html",
		"templates/"+vars["blogid"]+".html",
	)
	t.ExecuteTemplate(w, "base", "")
}
