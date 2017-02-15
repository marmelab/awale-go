package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", Awale)
	http.ListenAndServe(":8080", nil)
}

func Awale(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
