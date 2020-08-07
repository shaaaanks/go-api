package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shaaaanks/go-api/api/database"
	"github.com/shaaaanks/go-api/api/structs"
)

func GetEvents(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		events, err := db.FindAll()

		if err != nil {
			log.Fatalf("Error getting items from database: %v", err)
		}

		json.NewEncoder(w).Encode(events)
	}
}

func GetEvent(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		eventID := mux.Vars(r)["id"]

		event, err := db.Find(eventID)
		if err != nil {
			log.Fatalf("Error getting item from database: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(event)
	}
}

func CreateEvent(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newEvent structs.Event
		request, err := ioutil.ReadAll(r.Body)

		if err != nil {
			fmt.Fprintf(w, "Enter Event details")
		}

		json.Unmarshal(request, &newEvent)
		fmt.Printf("title: %v", newEvent.Title)
		fmt.Printf("description: %v", newEvent.Description)

		db.Create(newEvent)

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newEvent)
	}
}

func UpdateEvent(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		eventID := mux.Vars(r)["id"]
		var updatedEvent structs.Event

		request, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "Enter Event details")
		}

		json.Unmarshal(request, &updatedEvent)

		err = db.Update(eventID, updatedEvent)
		if err != nil {
			log.Fatalf("Error updating item in database: %v", err)
		}

		w.WriteHeader(http.StatusAccepted)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updatedEvent)

	}
}

func DeleteEvent(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		eventID := mux.Vars(r)["id"]

		err := db.Delete(eventID)
		if err != nil {
			log.Fatalf("Error deleting item from database: %v", err)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "The event with the ID %v has been deleted successfully", eventID)
	}
}
