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

	log.Fatal(http.ListenAndServe(":"+port, r))

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}
