package main

import (
	"log"
	"net/http"

	"github.com/bahattincinic/messenger-challenge/controllers"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc(
		"/auth/login",
		controllers.CreateAccessToken,
	).Methods(http.MethodPost)

	router.HandleFunc(
		"/auth/signup",
		controllers.Signup,
	).Methods(http.MethodPost)

	router.HandleFunc(
		"/users/",
		controllers.GetUserList,
	).Methods(http.MethodGet)

	// Start HTTP server.
	if err := http.ListenAndServe(":8090", router); err != nil {
		log.Fatal(err)
	}
}
