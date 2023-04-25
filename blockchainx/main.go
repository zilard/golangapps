package main

import (
	"time"

	"github.com/zilard/golangapps/blockchainx/network"
)

func main() {

	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func() {
		trRemote.SendMessage(trLocal.Addr(), []byte("Hello World"))
		time.Sleep(1 * time.Second)
	}()

	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)
	s.Start()

}
