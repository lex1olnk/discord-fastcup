package main

import (
    "fmt"
    "net/http"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", handleRoot)

    fmt.Println("Server listening to :8080 port")
    http.ListenAndServe(":8080", mux)
}

func handleRoot(
    w http.ResponseWriter, 
    r *http.Request,
) {
    fmt.Fprintf(w, "Hello world")
}
