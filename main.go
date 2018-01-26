package main

import (
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":"+port, nil)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>HELLO</h1>"))
}
