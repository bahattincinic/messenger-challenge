package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func dummyHandler(w http.ResponseWriter, r *http.Request) {}

func TestJsonResponseMiddleware(t *testing.T) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", dummyHandler).Methods("GET")

	router.Use(JSONResponseMiddleware)

	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	ctype := rr.Header().Get("Content-Type")

	if ctype != "application/json" {
		t.Errorf("content type header does not match: got %v want %v",
			ctype, "application/json")
	}
}
