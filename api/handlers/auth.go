package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bahattincinic/messenger-challenge/domain/models"
	"github.com/bahattincinic/messenger-challenge/domain/usecases"
)

// CreateAccessToken API Creates user Access token
func (h *BaseHandler) CreateAccessToken(w http.ResponseWriter, r *http.Request) {
	var login models.Login

	err := json.NewDecoder(r.Body).Decode(&login)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := usecases.CreateAccessToken(login, h.authRepo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, _ := json.Marshal(user)
	fmt.Fprintf(w, string(resp))
}

// Signup API Creates user
func (h *BaseHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var signup models.Signup

	err := json.NewDecoder(r.Body).Decode(&signup)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := usecases.CreateUser(signup, h.authRepo)
	resp, err := json.Marshal(user)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, string(resp))
}
