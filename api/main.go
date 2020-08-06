package main

import (
	"log"
	"net/http"

	"github.com/shaaaanks/go-api/api/database"

	"github.com/shaaaanks/go-api/api/router"
	kibisis "github.com/shaaaanks/go-api/kibisis"
)

func main() {
	var db database.DB
	var err error

	db.Driver, err = kibisis.GetDriver("arangoDB")
	if err != nil {
		log.Fatalf("Error loading database driver: %v", err)
	}

	err = db.Conn()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	err = db.Init()
	if err != nil {
		log.Fatalf("Error initialising database: %v", err)
	}

	r := router.GenerateRouter(&db)

	log.Fatal(http.ListenAndServe(":8080", r))
}
