package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	r.HandleFunc("/cal/{num1}/plus/{num2}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		num1Str := vars["num1"]
		num2Str := vars["num2"]

		num1, err := strconv.Atoi(num1Str)
		if err != nil {
			http.Error(w, "Invalid number", http.StatusBadRequest)
			return
		}

		num2, err := strconv.Atoi(num2Str)
		if err != nil {
			http.Error(w, "Invalid number", http.StatusBadRequest)
			return
		}

		total := num1 + num2
		fmt.Fprintf(w, "Total = %d\n", total)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "my name is surasak")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", r)
}
