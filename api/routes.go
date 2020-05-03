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
		http.HandlerFunc(handlers.Signup),
	},
	Route{
		"UserList",
		http.MethodGet,
		"/users/",
		handlers.GetUserList,
	},
	Route{
		"CurrentUser",
		http.MethodGet,
		"/me/",
		handlers.GetCurrentUser,
	},
	Route{
		"SendMessage",
		http.MethodPost,
		"/messages/{to}",
		handlers.CreateMessage,
	},
	Route{
		"ShowMessages",
		http.MethodGet,
		"/messages/{to}",
		handlers.GetMessages,
	},
}

// NewRouter returns Router instance
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Add common middlewares
	router.Use(middlewares.JSONResponseMiddleware)
	router.Use(middlewares.AuthenticationMiddleware)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			HandlerFunc(route.HandlerFunc)
	}

	return router
}
