package main

import (
	"context"
	"fmt"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

func database(name string) driver.Database {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
	})
	if err != nil {
		fmt.Errorf("Database connection error: %v", err)
	}

	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("root", "root"),
	})
	if err != nil {
		fmt.Errorf("Client error: %v", err)
	}

	db, err := client.Database(nil, name)
	if err != nil {
		fmt.Errorf("Database error: %v", err)
	}

	return db
}

type allevents []event

func getAllDocuments() allevents {
	db := database("events")

	ctx := context.Background()
	query := "FOR d IN events RETURN d"
	cursor, err := db.Query(ctx, query, nil)
	if err != nil {
		fmt.Print(err)
	}

	defer cursor.Close()
	var events = allevents{}

	for {
		var event event
		meta, err := cursor.ReadDocument(ctx, &event)
		event.Key = meta.Key

		events = append(events, event)
		if driver.IsNoMoreDocuments(err) {
			break
		}
	}

	return events
}

func collection(name string, db driver.Database) driver.Collection {
	col, err := db.Collection(nil, "events")
	if err != nil {
		fmt.Errorf("Collection error: %v", err)
	}

	return col
}
