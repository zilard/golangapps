// Route Between Nodes: 
// Given a directed graph, design an algorithm to find out 
// whether there is a route between two nodes

package main

import (
	"fmt"
	// "math"
	"sync"
	"reflect"
)



type Tree struct {
	Value int
	Visited bool
	Adjacents []*Tree
}



type queuenode struct {

	data interface{}
	next *queuenode

}



//------------------------------------------------------------------


// a go-routine safe FIFO data structure
type Queue struct {

	head *queuenode
	tail *queuenode
	count int
	lock *sync.Mutex

}




// creates a new pointe to a new queue
func New() *Queue {

	q := &Queue{}
	q.lock = &sync.Mutex{}
	return q

}


/*

Go's standard library provides mutual exclusion 
with sync.Mutex and its two methods:
- Lock
- Unlock

*/


func (q *Queue) Len() int {

	q.lock.Lock()
	defer q.lock.Unlock()
	return q.count

}



func (q *Queue) Push(item interface{}) {

	q.lock.Lock()
	defer q.lock.Unlock()

	n := &queuenode{ data: item }

	if q.tail == nil {
		q.tail = n
		q.head =n
	} else {
		q.tail.next = n
		q.tail = n
	}

	q.count++

}



func (q *Queue) Poll() interface{} {

	q.lock.Lock()
	defer q.lock.Unlock()

	if q.head == nil {
		return nil
	}

	n := q.head
	q.head = n.next

	if q.head == nil {
		q.tail = nil
	}
	q.count--

	return n.data

}



func (q *Queue) Peek() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()

	n := q.head
	if n == nil || n.data == nil {
		return nil
	}

	return n.data
}



func (q *Queue) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q == nil {
		return false
	}
	if q.count == 0 {
		return true
	}
	return false
}






func NewTree() *Tree {

	return &Tree{}

}








func isPathAvailable(t1 *Tree, t2 *Tree) bool {

	if t1 == nil || t2 == nil {
		return false
	}

	q := New()

	if t1.Visited == false {
		q.Push(t1)
		t1.Visited = true
		if t1.Value == t2.Value {
			return true
		}


	}


	for !q.IsEmpty() {

		node := q.Poll()

		// https://pkg.go.dev/reflect
		// go reflect package used to get the array of interface from the struct

		// reflect.ValueOf() Function in Golang is used to get the new Value 
		// initialized to the concrete value stored in the interface i.

		// ValueOf() returns a new Value initialized to the concrete value stored in the interface i

		// Elem() returns the value that the interface v contains or that the pointer v points to. 

		// FieldByName() returns the struct field with the given name

		// Interface() returns v's current value as an interface{}. 

		refValue := reflect.ValueOf(&node).Elem().FieldByName(string("Adjacents"))
		adjNodes := refValue.Interface().([]*Tree)

		for _, n := range adjNodes {
			if n.Visited == false {
				if n.Value == t2.Value {
					return true
				}
				n.Visited = true
				q.Push(n)
			}
		}

	}

	return false

}




func main() {

	t1 := &Tree{
		Value: 1,
		}

        t2 := &Tree{
                Value: 2,
                }

	t1.Adjacents = append(t1.Adjacents, t2)


	fmt.Printf("t1: %v\n", t1)

        pathAvailable := isPathAvailable(t1, t2)

        fmt.Printf("Is Path Available t1 -> t2 ? %v\n", pathAvailable)


}



