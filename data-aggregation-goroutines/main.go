package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	userName := fetchUser()

	// will buffer my channel
	// because unbuffered channel will always block
	// until somebody is reading from it
	respch := make(chan any, 2)

	go fetchUserLikes(userName, respch)
	go fetchUserMatch(userName, respch)

	// in golang you can range over channels
	for resp := range respch {
		fmt.Println("resp: ", resp)
	}

	fmt.Println("took: ", time.Since(start))
}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)

	return "BOB"
}

func fetchUserLikes(userName string, respch chan any) {
	time.Sleep(time.Millisecond * 150)

	respch <- 11
}

func fetchUserMatch(userName string, respch chan any) {
	time.Sleep(time.Millisecond * 100)

	respch <- "ANNA"
}
