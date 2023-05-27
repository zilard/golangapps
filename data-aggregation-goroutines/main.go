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

	// we need to close the response channel, so that
	// we dont have deadlock when we try to range over channel and wait for data
	close(respch)

	// in golang you can range over channels
	// we gonna have a deadlock if we dont close the channel
	for resp := range respch {
		fmt.Println("resp: ", resp)

		/*
			// casting channel from any to a specific type
			likes, ok := resp.(int)
			// if ok then you know that's an int
			if ok {
				fmt.Println("likes: ", likes)
			}
		*/

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
