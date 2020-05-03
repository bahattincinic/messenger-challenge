package api

import (
	"net/http"

	"github.com/bahattincinic/messenger-challenge/api/handlers"
	"github.com/bahattincinic/messenger-challenge/api/middlewares"
	"github.com/gorilla/mux"
)

// Route is a API Route Object
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.Handler
}

// Routes is a list of API Route
type Routes []Route

var routes = Routes{
	Route{
		"Login",
		http.MethodPost,
		"/auth/login",
		http.HandlerFunc(handlers.CreateAccessToken),
	},
	Route{
		"Signup",
		http.MethodPost,
		"/auth/signup",
		http.HandlerFunc(handlers.Signup),
	},
	Route{
		"UserList",
		http.MethodGet,
		"/users/",
		middlewares.AuthenticationMiddleware(
			http.HandlerFunc(handlers.GetUserList),
		),
	},
	Route{
		"CurrentUser",
		http.MethodGet,
		"/me/",
		middlewares.AuthenticationMiddleware(
			http.HandlerFunc(handlers.GetCurrentUser),
		),
	},
	Route{
		"SendMessage",
		http.MethodPost,
		"/messages/{to}",
		middlewares.AuthenticationMiddleware(
			http.HandlerFunc(handlers.CreateMessage),
		),
	},
	Route{
		"ShowMessages",
		http.MethodGet,
		"/messages/{to}",
		middlewares.AuthenticationMiddleware(
			http.HandlerFunc(handlers.GetMessages),
		),
	},
}

// NewRouter returns Router instance
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Add common middlewares
	router.Use(middlewares.JSONResponseMiddleware)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
