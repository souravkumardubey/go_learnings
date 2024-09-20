package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    if r.Method != "GET" {
        w.Header().Set("Allow", "GET")
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }
	
    fmt.Fprintf(w, "Hello, World!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "About Page")
}

func main() {
    http.HandleFunc("/", helloHandler)
    http.HandleFunc("/about", aboutHandler)

    fmt.Println("Starting server at port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
