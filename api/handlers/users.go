package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bahattincinic/messenger-challenge/domain/models"
	"github.com/bahattincinic/messenger-challenge/domain/usecases"
)

// GetUserList handler returns list of user
func GetUserList(w http.ResponseWriter, r *http.Request) {
	users := usecases.GetUsers()
	resp, _ := json.Marshal(users)

	fmt.Fprintf(w, string(resp))
}

// GetCurrentUser handler returns authanticated user informations
func GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(models.User)
	resp, _ := json.Marshal(user)
	fmt.Fprintf(w, string(resp))
}
