package main

import (
	"fmt"
)


type Queue struct {
	count int
	lock *int
}


func (q *Queue) doSomething() {

	fmt.Printf("count: %v\n", q.count)
        fmt.Printf("lock: %v\n", *(q.lock))

}


func main() {

	l := 5
	q := Queue{
		count: 6,
		lock: &l,
		}

	q.doSomething()

}



