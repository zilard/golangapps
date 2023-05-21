package main

import "net/http"

func main() {

	http.HandleFunc("/user", handleGetUserByID)

	http.ListenAndServe(":3000", nil)
}

func handleGetUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte("not allowed"))
		return
	}
}
