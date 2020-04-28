package controllers

import (
	"fmt"
	"net/http"
)

// CreateAccessToken API Creates user Access token
func CreateAccessToken(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create Access Token\n")
}
