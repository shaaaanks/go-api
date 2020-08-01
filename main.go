package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Learn the Go",
	},
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World")
}

func getEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}

func getEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, event := range events {
		if event.ID == eventID {
			json.NewEncoder(w).Encode(event)
		}
	}
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	request, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Enter Event details")
	}

	json.Unmarshal(request, &newEvent)
	events = append(events, newEvent)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newEvent)
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var updatedEvent event

	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter Event details")
	}

	json.Unmarshal(request, &updatedEvent)

	for i, event := range events {
		if event.ID == eventID {
			event.Title = updatedEvent.Title
			event.Description = updatedEvent.Description

			events = append(events[:i], event)

			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(event)
		}
	}
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for i, event := range events {
		if event.ID == eventID {
			events = append(events[:i], events[i+1:]...)

			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "The event with the ID %v has been deleted successfully", eventID)
		}
	}
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
