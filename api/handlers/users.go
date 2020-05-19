package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bahattincinic/messenger-challenge/domain/usecases"
)

// GetUserList handler returns list of user
func (h *BaseHandler) GetUserList(w http.ResponseWriter, r *http.Request) {
	users := usecases.GetUsers(h.userRepo)
	resp, _ := json.Marshal(users)

	fmt.Fprint(w, string(resp))
}

// GetCurrentUser handler returns authanticated user informations
func (h *BaseHandler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	user, err := GetUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, _ := json.Marshal(user)
	fmt.Fprint(w, string(resp))
}
