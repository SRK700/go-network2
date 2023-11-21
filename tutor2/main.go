package main

import (
	"fmt"
	"net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/" {
            fmt.Fprint(w, "Hello, World!\n")
        } else if r.URL.Path == "/about" {
            fmt.Fprint(w, "My name is surasak\n")
        } else {
            http.NotFound(w, r)
            return
        }
    })

    http.ListenAndServe(":8080", nil)
}
