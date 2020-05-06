package api

import (
	"net/http"

	"github.com/bahattincinic/messenger-challenge/api/handlers"
	"github.com/bahattincinic/messenger-challenge/api/middlewares"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
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

func getRoutes(handler *handlers.BaseHandler) Routes {
	return Routes{
		Route{
			"Login",
			http.MethodPost,
			"/auth/login",
			handler.CreateAccessToken,
		},
		Route{
			"Signup",
			http.MethodPost,
			"/auth/signup",
			handler.Signup,
		},
		Route{
			"UserList",
			http.MethodGet,
			"/users/",
			handler.GetUserList,
		},
		Route{
			"CurrentUser",
			http.MethodGet,
			"/me/",
			handler.GetCurrentUser,
		},
		Route{
			"SendMessage",
			http.MethodPost,
			"/messages/{to}",
			handler.CreateMessage,
		},
		Route{
			"ShowMessages",
			http.MethodGet,
			"/messages/{to}",
			handler.GetMessages,
		},
	}
}

// NewRouter returns Router instance
func NewRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	handler := handlers.NewBaseHandler(db)
	routes := getRoutes(handler)

	// Add common middlewares
	router.Use(middlewares.JSONResponseMiddleware)
	router.Use(middlewares.AuthenticationMiddleware(db))

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			HandlerFunc(route.HandlerFunc)
	}

	return router
}
