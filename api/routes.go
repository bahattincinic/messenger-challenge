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
	HandlerFunc func(http.ResponseWriter, *http.Request)
}

// Routes is a list of API Route
type Routes []Route

var routes = Routes{
	Route{
		"Login",
		http.MethodPost,
		"/auth/login",
		handlers.CreateAccessToken,
	},
	Route{
		"Signup",
		http.MethodPost,
		"/auth/signup",
		handlers.Signup,
	},
	Route{
		"UserList",
		http.MethodGet,
		"/users/",
		middlewares.AuthenticationMiddleware(handlers.GetUserList),
	},
	Route{
		"CurrentUser",
		http.MethodGet,
		"/me/",
		middlewares.AuthenticationMiddleware(handlers.GetCurrentUser),
	},
	Route{
		"SendMessage",
		http.MethodPost,
		"/messages/{to}",
		middlewares.AuthenticationMiddleware(handlers.CreateMessage),
	},
	Route{
		"ShowMessages",
		http.MethodGet,
		"/messages/{to}",
		middlewares.AuthenticationMiddleware(handlers.GetMessages),
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
			HandlerFunc(route.HandlerFunc)
	}

	return router
}
