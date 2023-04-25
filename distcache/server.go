package main

import "distcache/cache"

type ServerOpts struct {
	ListenAddr string
	IsLeader   bool
}

type Server struct {
	ServerOpts

	cache cache.Cache
}

func NewServer(opts ServerOpts, c cache.Cacher) *Server {

}
