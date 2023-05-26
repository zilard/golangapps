package main

import "fmt"

type Opts struct {
	maxConn int
	id      string
	tls     bool
}

type Server struct {
	Opts
}

func newServer(opts Opts) *Server {
	return &Server{
		Opts: opts,
	}
}

func main() {
	s := newServer(Opts{})
	fmt.Printf("%+v\n", s)
}
