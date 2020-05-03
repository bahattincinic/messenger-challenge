package middlewares

import (
	"context"
	"net/http"

	"github.com/bahattincinic/messenger-challenge/domain/usecases"
)

const accessTokenHeader = "X-Access-Token"

// UserCtxKey is a key name of user context
const UserCtxKey = "user"

// AuthenticationMiddleware user access token
func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get(accessTokenHeader)

		if len(accessToken) == 0 {
			next.ServeHTTP(w, r)
			return
		}

		user, err := usecases.CheckAccessToken(accessToken)

		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(
			r.Context(), UserCtxKey, user,
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
