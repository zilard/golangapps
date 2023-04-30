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

func handler(w http.ResponseWriter, r *http.Request) {
	// where is my DB
}

func main() {
	http.HandleFunc("/", handler)

	s := &Store{}
	Execute(myExecuteFunc(s))
}

// type function
// this is coming from a third party lib
type ExecuteFn func(string)

func Execute(fn ExecuteFn) {
	fn("FOO BAR BAZ")
}
