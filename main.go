package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bahattincinic/messenger-challenge/controllers"
	"github.com/bahattincinic/messenger-challenge/middlewares"
	"github.com/gorilla/handlers"
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
		middlewares.AuthenticationMiddleware(controllers.GetUserList),
	).Methods(http.MethodGet)

	router.HandleFunc(
		"/me/",
		middlewares.AuthenticationMiddleware(controllers.GetCurrentUser),
	).Methods(http.MethodGet)

	// Start HTTP server.
	err := http.ListenAndServe(
		":8090",
		handlers.LoggingHandler(os.Stdout, router),
	)

	if err != nil {
		log.Fatal(err)
	}
}
