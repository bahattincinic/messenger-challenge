package middlewares

import (
	"net/http"

	"github.com/bahattincinic/messenger-challenge/usecases"
)

// MiddlewareCallback is a callback interface
type MiddlewareCallback func(w http.ResponseWriter, r *http.Request)

// AuthenticationMiddleware user access token
func AuthenticationMiddleware(next MiddlewareCallback) MiddlewareCallback {
	return func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("X-Access-Token")

		if len(accessToken) == 0 {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		isExists := usecases.CheckAccessToken(accessToken)

		if isExists == false {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}
