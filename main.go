package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)

	log.Fatal(http.ListenAndServe(":8080", router))
}
