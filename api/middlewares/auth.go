package middlewares

import (
	"context"
	"net/http"

	"github.com/bahattincinic/messenger-challenge/domain/repositories"
	"github.com/bahattincinic/messenger-challenge/domain/usecases"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

const accessTokenHeader = "X-Access-Token"

// UserCtxKey is a key name of user context
const UserCtxKey = "user"

// AuthenticationMiddleware user access token
func AuthenticationMiddleware(db *gorm.DB) mux.MiddlewareFunc {
	authRepo := repositories.NewAuthRepo(db)
	userRepo := repositories.NewUserRepo(db)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			accessToken := r.Header.Get(accessTokenHeader)

			if len(accessToken) == 0 {
				next.ServeHTTP(w, r)
				return
			}

			user, err := usecases.CheckAccessToken(
				accessToken, authRepo, userRepo,
			)

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
}
