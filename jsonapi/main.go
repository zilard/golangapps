package main

import (
	"encoding/json"
	"net/http"
)

type apiError struct {
	err error
	status int
}

// implements the Error interface
func (e apiError) Error() string {
	return e.err
}


func main() {

	http.HandleFunc("/user", handleGetUserByID)

	http.ListenAndServe(":3000", nil)
}

func handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		if err := writeJSON(w, http.StatusMethodNotAllowed, "any")
		return

	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)

}
