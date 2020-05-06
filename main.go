package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bahattincinic/messenger-challenge/api"
	"github.com/bahattincinic/messenger-challenge/domain/models"
	"github.com/gorilla/handlers"
)

func main() {
	db, dbErr := models.InitDatabase()
	if dbErr != nil {
		log.Fatalf("failed to connect database: %s\n", dbErr)
	}

	router := api.NewRouter(db)

	// Start HTTP server.
	err := http.ListenAndServe(
		":8090",
		handlers.LoggingHandler(os.Stdout, router),
	)
	log.Fatal(err)
}
