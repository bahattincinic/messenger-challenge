package handlers

import (
	"errors"
	"net/http"

	"github.com/bahattincinic/messenger-challenge/api/middlewares"
	"github.com/bahattincinic/messenger-challenge/domain/models"
)

// GetUser returns current user
func GetUser(r *http.Request) (user models.User, err error) {
	user, found := r.Context().Value(middlewares.UserCtxKey).(models.User)

	if found == false {
		err = errors.New("Forbidden")
	}

	return
}
