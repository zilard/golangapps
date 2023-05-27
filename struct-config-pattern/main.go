package main

import "fmt"

type OptFunc func(*Opts)

type Opts struct {
	maxConn int
	id      string
	tls     bool
}

func defaultOpts() Opts {
	return Opts{
		maxConn: 10,
		id:      "default",
		tls:     false,
	}
}

func withID(id string) OptFunc {
	return func(opts *Opts) {
		opts.id = id
	}
}

func withTLS(opts *Opts) {
	opts.tls = true
}

func withMaxConn(n int) OptFunc {
	return func(opts *Opts) {
		opts.maxConn = n
	}
}

type Server struct {
	Opts
}

func newServer(opts ...OptFunc) *Server {
	o := defaultOpts()
	for _, fn := range opts {
		fn(&o)
	}
	return &Server{
		Opts: o,
	}
}

func main() {
	s := newServer(withMaxConn(99), withID("bar foo waz"))
	fmt.Printf("%+v\n", s)
}
