package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/api/", ProvideStatusCode)
	http.ListenAndServe(":8080", nil)
}

func ProvideStatusCode(w http.ResponseWriter, r *http.Request) {
	str := strings.TrimPrefix(r.URL.Path, "/api/")
	code, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(code)
}
