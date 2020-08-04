package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	kibisis "github.com/shaaaanks/go-api"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World")
}

func getEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	driver, err := kibisis.GetDriver("arangoDB")
	if err != nil {
		log.Fatalf("Error loading database driver: %v", err)
	}

	err = driver.Conn()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	err = driver.Init()
	if err != nil {
		log.Fatalf("Error initialising database: %v", err)
	}

	events, err := driver.FindAll()
	if err != nil {
		log.Fatalf("Error getting items from database: %v", err)
	}

	json.NewEncoder(w).Encode(events)
}

func getEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	driver, err := kibisis.GetDriver("arangoDB")
	if err != nil {
		log.Fatalf("Error loading database driver: %v", err)
	}

	err = driver.Conn()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	err = driver.Init()
	if err != nil {
		log.Fatalf("Error initialising database: %v", err)
	}

	event, err := driver.Find(eventID)
	if err != nil {
		log.Fatalf("Error getting item from database: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	driver, err := kibisis.GetDriver("arangoDB")
	if err != nil {
		log.Fatalf("Error loading database driver: %v", err)
	}

	err = driver.Conn()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	err = driver.Init()
	if err != nil {
		log.Fatalf("Error initialising database: %v", err)
	}

	var newEvent Event
	request, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Enter Event details")
	}

	json.Unmarshal(request, &newEvent)
	fmt.Printf("title: %v", newEvent.Title)
	fmt.Printf("description: %v", newEvent.Description)

	driver.Create(newEvent)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newEvent)
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var updatedEvent Event

	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter Event details")
	}

	json.Unmarshal(request, &updatedEvent)

	driver, err := kibisis.GetDriver("arangoDB")
	if err != nil {
		log.Fatalf("Error loading database driver: %v", err)
	}

	err = driver.Conn()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	err = driver.Init()
	if err != nil {
		log.Fatalf("Error initialising database: %v", err)
	}

	err = driver.Update(eventID, updatedEvent)
	if err != nil {
		log.Fatalf("Error updating item in database: %v", err)
	}

	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedEvent)

}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	driver, err := kibisis.GetDriver("arangoDB")
	if err != nil {
		log.Fatalf("Error loading database driver: %v", err)
	}

	err = driver.Conn()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	err = driver.Init()
	if err != nil {
		log.Fatalf("Error initialising database: %v", err)
	}

	err = driver.Delete(eventID)
	if err != nil {
		log.Fatalf("Error deleting item from database: %v", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "The event with the ID %v has been deleted successfully", eventID)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/event", createEvent).Methods("POST")
	router.HandleFunc("/events", getEvents).Methods("GET")
	router.HandleFunc("/event/{id}", getEvent).Methods("GET")
	router.HandleFunc("/event/{id}", updateEvent).Methods("PATCH")
	router.HandleFunc("/event/{id}", deleteEvent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
