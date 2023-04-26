package main

import (
	"log"
	"net"
	"time"

	"github.com/zilard/golangapps/distcache/cache"
)

func main() {
	opts := ServerOpts{
		ListenAddr: ":3000",
		IsLeader:   true,
	}

	go func() {
		time.Sleep(time.Second * 2)
		conn, err := net.Dial("tcp", ":3000")
		if err != nil {
			log.Fatal(err)
		}

		conn.Write([]byte("SET Foo Bar 2500000000000"))

		time.Sleep(time.Second * 2)
		conn.Write([]byte("GET Foo"))
	}()

	server := NewServer(opts, cache.New())
	server.Start()
}
