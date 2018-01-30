package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// base template
var hometemplate *template.Template

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000" // gin port
	}

	// cache base template
	hometemplate = template.Must(template.ParseFiles("./templates/base.html",
		"./templates/home.html"))

	// Mux Router
	r := mux.NewRouter()

	// serve static files
	r.PathPrefix("/static/").
		Handler(http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	// routes
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/blog/{blogid}", blogHandler)

	// serve
	log.Fatal(http.ListenAndServe(":"+port, r))

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	hometemplate.ExecuteTemplate(w, "base", "")
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["blogid"]

	t, err := template.ParseFiles("./templates/base.html",
		"./templates/blogs/blog"+id+".html")

	if err != nil {
		// go home
		http.Redirect(w, r, "/", 301)
		return
	}

	t.ExecuteTemplate(w, "base", "")
}
