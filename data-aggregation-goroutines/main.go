package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	userName := fetchUser()
	likes := fetchUserLikes(userName)
	match := fetchUserMatch(userName)

	fmt.Println("likes: ", likes)
	fmt.Println("match: ", match)
	fmt.Println("took: ", time.Since(start))
}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)

	return "BOB"
}

func fetchUserLikes(userName string) int {
	time.Sleep(time.Millisecond * 150)

	return 11
}

func fetchUserMatch(userName string) string {
	time.Sleep(time.Millisecond * 100)

	return "ANNA"
}
