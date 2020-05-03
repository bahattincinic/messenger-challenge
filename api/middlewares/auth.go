package middlewares

import (
	"context"
	"net/http"

	"github.com/bahattincinic/messenger-challenge/domain/usecases"
)

// AuthenticationMiddleware user access token
func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		ctx := context.WithValue(r.Context(), "user", user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
