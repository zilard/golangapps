package main

import (
	"encoding/json"
	"net/http"
)

func main() {

	http.HandleFunc("/user", handleGetUserByID)

	http.ListenAndServe(":3000", nil)
}

func handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		if err := writeJSON(w, http.StatusMethodNotAllowed, "any"); err != nil {
			return
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)

}
