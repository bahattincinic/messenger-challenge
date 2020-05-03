package handlers

import (
	"errors"
	"net/http"

	"github.com/bahattincinic/messenger-challenge/domain/models"
)

// GetUser returns current user
func GetUser(r *http.Request) (user models.User, err error) {
	user, found := r.Context().Value("user").(models.User)

	if found == false {
		err = errors.New("Forbidden")
	}

	return
}
