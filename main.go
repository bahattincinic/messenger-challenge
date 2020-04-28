package main

import (
	"log"
	"net/http"

	"github.com/bahattincinic/messenger-challenge/controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc(
		"/auth/access-token",
		controllers.CreateAccessToken,
	).Methods(http.MethodPost)

	// Start HTTP server.
	if err := http.ListenAndServe(":8090", router); err != nil {
		log.Fatal(err)
	}
}
