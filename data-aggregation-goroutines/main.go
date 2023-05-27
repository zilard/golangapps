package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	userName := fetchUser() // takes 100ms to execute, delay added

	// will buffer my channel
	// because unbuffered channel will always block
	// until somebody is reading from it
	respch := make(chan any, 2)

	// use async wait group
	wg := &sync.WaitGroup{}

	// we have to workers that needs to communicate with the wait group
	// telling it that they are done
	wg.Add(2)
	// each time we call wg.Done() it's going to substract from this number
	// and when it goes to zero then wg.Wait() will unblock

	// then we gonna do our work
	// we do the wg.Done() in our asynchronous functions
	// we need to provide the waitgroup as parameter to these functions
	go fetchUserLikes(userName, respch, wg)
	go fetchUserMatch(userName, respch, wg)

	// right before we close the channel we call wg wait
	// this is going to block until we have 2 wg.Done()
	wg.Wait()

	// we need to close the response channel, so that
	// we dont have deadlock when we try to range over channel and wait for data
	close(respch)

	// we are closing the channel who guarantees that these guys
	// fetchUserLikes and fetchUserMatch are already done

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

func fetchUserLikes(userName string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150)

	respch <- 11

	// we do the wg.Done() in our asynchronous functions
	wg.Done()
}

func fetchUserMatch(userName string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)

	respch <- "ANNA"

	// we do the wg.Done() in our asynchronous functions
	wg.Done()
}
