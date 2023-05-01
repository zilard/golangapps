package main

import (
	"fmt"
	"net/http"
)

type DB interface {
	Store(string) error
}

type Store struct{}

func (s *Store) Store(value string) error {
	fmt.Println("storing into db", value)
	return nil
}

// we are executing the function and we are
// also storing into the DB
func myExecuteFunc(db DB) ExecuteFn {
	return func(s string) {
		fmt.Println("my ex func", s)
		db.Store(s)
	}
}

func makeHTTPFunc(db DB, fn httpFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(db, w, r); err != nil {
			//
		}
	}
}

func main() {
	s := &Store{}
	http.HandleFunc("/", makeHTTPFunc(s, handler))
	Execute(myExecuteFunc(s))
}

func handler(db DB, w http.ResponseWriter, r *http.Request) error {
	return nil
}

type httpFunc func(db DB, w http.ResponseWriter, r *http.Request) error

// type function
// this is coming from a third party lib
type ExecuteFn func(string)

func Execute(fn ExecuteFn) {
	fn("FOO BAR BAZ")
}
