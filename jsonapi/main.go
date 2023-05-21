package main

import (
	"encoding/json"
	"net/http"
)

var ErrUserInvalid = apiError{Err: "user not valid",
	Status: http.StatusForbidden}

type apiError struct {
	Err    string
	Status int
}

// implements the Error interface
func (e apiError) Error() string {
	return e.Err
}

type apiFunc func(http.ResponseWriter, *http.Request) error

// we have a centralized HTTP Handler
func makeHTTPHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			if e, ok := err.(apiError); ok {
				writeJSON(w, e.Status, e)
				return
			}
			writeJSON(w, http.StatusInternalServerError,
				apiError{Err: "internal server",
					Status: http.StatusInternalServerError})
		}
	}
}

func main() {

	http.HandleFunc("/user", makeHTTPHandler(handleGetUserByID))

	http.ListenAndServe(":3000", nil)
}

type User struct {
	ID    int
	Valid bool
}

func handleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return apiError{Err: "invalid method", Status: http.StatusMethodNotAllowed}
	}

	user := User{}
	if !user.Valid {
		return ErrUserInvalid
	}

	return writeJSON(w, http.StatusOK, user)
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)

}
