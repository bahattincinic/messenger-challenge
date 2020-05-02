package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bahattincinic/messenger-challenge/api"
	"github.com/gorilla/handlers"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := api.NewRouter()

	// Start HTTP server.
	err := http.ListenAndServe(
		":8090",
		handlers.LoggingHandler(os.Stdout, router),
	)
	log.Fatal(err)
}
