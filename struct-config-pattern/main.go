package main

import "fmt"

type Server struct {
	maxConn int
	id      string
	tls     bool
}

func newServer(maxConn int, id string, tls bool) *Server {
	return &Server{
		maxConn: maxConn,
		id:      id,
		tls:     tls,
	}
}

func main() {
	s := newServer(1, "foo", false)
	fmt.Printf("%+v\n", s)
}
