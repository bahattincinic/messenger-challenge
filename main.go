package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bahattincinic/messenger-challenge/api"
	"github.com/bahattincinic/messenger-challenge/config"
	"github.com/bahattincinic/messenger-challenge/domain/models"
	"github.com/gorilla/handlers"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)

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
