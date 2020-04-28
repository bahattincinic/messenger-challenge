package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bahattincinic/messenger-challenge/usecases"
)

func GetUserList(w http.ResponseWriter, r *http.Request) {
	users := usecases.GetUsers()
	resp, _ := json.Marshal(users)

	fmt.Fprintf(w, string(resp))
}
