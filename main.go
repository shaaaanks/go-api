package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type event struct {
	Key         string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World")
}

func getEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	events := getAllDocuments()

	json.NewEncoder(w).Encode(events)
}

func getEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	database := database("events")
	collection := collection("events", database)
	ctx := context.Background()
	var event event
	collection.ReadDocument(ctx, eventID, &event)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	request, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Enter Event details")
	}

	json.Unmarshal(request, &newEvent)
	fmt.Printf("title: %v", newEvent.Title)
	fmt.Printf("description: %v", newEvent.Description)

	if newEvent.Title != "" {
		database := database("events")
		collection := collection("events", database)

		meta, err := collection.CreateDocument(nil, newEvent)
		if err != nil {
			fmt.Errorf("Creation error: %v", err)
		}
		log.Println(meta)

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newEvent)
	}
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var updatedEvent event

	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter Event details")
	}

	json.Unmarshal(request, &updatedEvent)
	// fmt.Print(updatedEvent.)

	ctx := context.Background()
	database := database("events")
	collection := collection("events", database)
	meta, err := collection.UpdateDocument(ctx, eventID, updatedEvent)
	if err != nil {
		fmt.Errorf("Update error: %v", err)
	}

	fmt.Print(meta)

	w.WriteHeader(http.StatusAccepted)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(meta)

}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	ctx := context.Background()
	database := database("events")
	collection := collection("events", database)
	meta, err := collection.RemoveDocument(ctx, eventID)
	if err != nil {
		fmt.Errorf("Deletion error: %v", err)
	}

	fmt.Print(meta)

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
