package middlewares

import (
	"net/http"

	"github.com/bahattincinic/messenger-challenge/models"
	"github.com/bahattincinic/messenger-challenge/usecases"
)

// WrappedMiddlewareFunc is a wrapped function reference
type WrappedMiddlewareFunc func(w http.ResponseWriter, r *http.Request)

// RequestMiddlewareCallback is a callback function
type RequestMiddlewareCallback func(w http.ResponseWriter, r *http.Request, user models.User)

// AuthenticationMiddleware user access token
func AuthenticationMiddleware(next RequestMiddlewareCallback) WrappedMiddlewareFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("X-Access-Token")

		if len(accessToken) == 0 {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		user, err := usecases.CheckAccessToken(accessToken)

		if err != nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next(w, r, user)
	}
}
