package main

import (
	"fastcup/api"
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handleRoot)

	http.HandleFunc("/match/", func(w http.ResponseWriter, r *http.Request) {
		api.MatchHandler(w, r)
	})

	http.HandleFunc("/matches", func(w http.ResponseWriter, r *http.Request) {
		api.MatchesHandler(w, r)
	})

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func handleRoot(
	w http.ResponseWriter,
	r *http.Request,
) {
	fmt.Fprintf(w, "Hello world")
}
