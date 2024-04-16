package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Your solution goes here. Good luck!

	http.HandleFunc("/hello", yourFunction)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func yourFunction(w http.ResponseWriter, r *http.Request) {
	name:= r.URL.Query().Get("name")
	if r.URL.Query().Get("name") == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello, %s", name)
}
