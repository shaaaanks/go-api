package main

import (
	"fmt"
	"net/http"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "got"}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "posted"}`))
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "put"}`))
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "deleted"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "404"}`))
	}
}

func main() {
	// s := &server{}
	// http.Handle("/", s)
	// log.Fatal(http.ListenAndServe(":8080", nil))

	a := 4
	b := &a

	fmt.Printf("Address of var a: %p\n", b)
	fmt.Printf("Value of var a: %v\n", *b)

	*b = 2
	fmt.Printf("Address of var a: %p\n", b)
	fmt.Printf("Value of var a: %v\n", *b)

	a = 1
	fmt.Printf("Address of var a: %p\n", b)
	fmt.Printf("Value of var a: %v\n", *b)

}
