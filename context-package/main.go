package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()

	// you can store variables inside context
	ctx := context.WithValue(context.Background(), "foo", "bar")

	// ctx := context.Background()
	userID := 10
	val, err := fetchUserData(ctx, userID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result: ", val)
	fmt.Println("took: ", time.Since(start))

}

type Response struct {
	value int
	err   error
}

func fetchUserData(ctx context.Context, userID int) (int, error) {

	// fetch variable from context
	val := ctx.Value("foo")
	fmt.Println(val.(string))

	// WithTimeout - we gonna create some context that is
	// going to timeout after a certain timeframe
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()
	respch := make(chan Response)

	// we schedule this guy into a goroutine
	go func() {
		val, err := fetchThirdPartyStuffWhichCanBeSlow()
		respch <- Response{
			value: val,
			err:   err,
		}
	}()

	// if we call fetch user data
	// after 200 milliseconds this channel will be closed
	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("fetching data from third party took too long")
		case resp := <-respch:
			return resp.value, resp.err
		}
	}
}

func fetchThirdPartyStuffWhichCanBeSlow() (int, error) {
	time.Sleep(time.Millisecond * 150)
	return 777, nil
}
