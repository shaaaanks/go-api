package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shaaaanks/go-api/api/database"
	"github.com/shaaaanks/go-api/api/handlers"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World")
}

// GenerateRouter - Adds each route to a router and returns it
func GenerateRouter(db *database.DB) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index).Methods("GET")

	router.HandleFunc("/event", handlers.CreateEvent(db)).Methods("POST")
	router.HandleFunc("/events", handlers.GetEvents(db)).Methods("GET")
	router.HandleFunc("/event/{id}", handlers.GetEvent(db)).Methods("GET")
	router.HandleFunc("/event/{id}", handlers.UpdateEvent(db)).Methods("PATCH")
	router.HandleFunc("/event/{id}", handlers.DeleteEvent(db)).Methods("DELETE")

	return router
}
