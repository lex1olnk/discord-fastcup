package main

import (
    "fmt"
    "net/http"
    "fastcup/api"
)

func main() {
    http.HandleFunc("/", handleRoot)

    http.HandleFunc("/match/", func(w http.ResponseWriter, r *http.Request) {
        api.MatchHandler(w, r)
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
