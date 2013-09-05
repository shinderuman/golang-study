package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "root")
	})
	http.HandleFunc("/aaaa", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "aaaa")
	})
	http.ListenAndServe(":8080", nil)
}
