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

// func getEvent(w http.ResponseWriter, r *http.Request) {
// 	eventID := mux.Vars(r)["id"]

// 	for _, event := range events {
// 		if event.ID == eventID {
// 			w.Header().Set("Content-Type", "application/json")
// 			res, _ := json.Marshal(event)
// 			w.Write(res)
// 		}
// 	}
// }

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

// func updateEvent(w http.ResponseWriter, r *http.Request) {
// 	eventID := mux.Vars(r)["id"]
// 	var updatedEvent event

// 	request, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Enter Event details")
// 	}

// 	json.Unmarshal(request, &updatedEvent)

// 	for i, event := range events {
// 		if event.ID == eventID {
// 			event.Title = updatedEvent.Title
// 			event.Description = updatedEvent.Description

// 			events = append(events[:i], event)

// 			w.WriteHeader(http.StatusAccepted)
// 			w.Header().Set("Content-Type", "application/json")
// 			res, _ := json.Marshal(event)
// 			w.Write(res)
// 		}
// 	}
// }

// func deleteEvent(w http.ResponseWriter, r *http.Request) {
// 	eventID := mux.Vars(r)["id"]

// 	for i, event := range events {
// 		if event.ID == eventID {
// 			events = append(events[:i], events[i+1:]...)

// 			w.WriteHeader(http.StatusOK)
// 			w.Header().Set("Content-Type", "application/json")
// 			fmt.Fprintf(w, "The event with the ID %v has been deleted successfully", eventID)
// 		}
// 	}
// }

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/event", createEvent).Methods("POST")
	router.HandleFunc("/events", getEvents).Methods("GET")
	// router.HandleFunc("/event/{id}", getEvent).Methods("GET")
	// router.HandleFunc("/event/{id}", updateEvent).Methods("PATCH")
	// router.HandleFunc("/event/{id}", deleteEvent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
