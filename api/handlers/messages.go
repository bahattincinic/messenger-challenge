package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bahattincinic/messenger-challenge/domain/models"
	"github.com/bahattincinic/messenger-challenge/domain/usecases"
	"github.com/gorilla/mux"
)

// GetMessages returns user messages
func GetMessages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var toUser string = vars["to"]

	user, err := GetUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	messages, err := usecases.GetUserMessages(user, toUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, _ := json.Marshal(messages)
	fmt.Fprintf(w, string(resp))
}

// CreateMessage creates message
func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var message models.MessageCreate
	vars := mux.Vars(r)
	var toUser string = vars["to"]

	user, authErr := GetUser(r)
	if authErr != nil {
		http.Error(w, authErr.Error(), http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&message)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdMessage, err := usecases.CreateMessage(user, toUser, message)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, _ := json.Marshal(createdMessage)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, string(resp))
}
