
package main

import (
	"fmt"
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



type Queue struct {
	head *queuenode
	tail *queuenode
	count int
	lock sync.Mutex
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




func main() {

	q := &Queue{}

	t1 := &Tree{
		Value: 1,
		}

	q.Push(t1)


	node := q.Poll()

	refVal := reflect.ValueOf(&node)

	fmt.Printf("refVal type: %v\n", reflect.TypeOf(refVal))

	// refVal type: reflect.Value
        // Value is the reflection interface to a Go value.

	// type Value struct {
		// contains filtered or unexported fields
	// }


	// func (v Value) Elem() Value
	// Elem returns the value that the interface v contains 
	// or that the pointer v points to.



// node:  queuenode.data
// queuenode.data = &Tree{..}

//      func ValueOf(i any) Value
//      ValueOf returns a new Value initialized to the concrete value stored in the interface i
        val := reflect.ValueOf(&node)
        fmt.Printf("val: %v\n", val)
        fmt.Printf("val type: %v\n", reflect.TypeOf(val))
/*
	val: 0xc000010250      <------------ &queuenode.data
	val type: reflect.Value

*/



//      func (v Value) Elem() Value
//      Elem returns the value that the interface v contains or that the pointer v points to
	val1 := reflect.ValueOf(&node).Elem()
        fmt.Printf("val1: %v\n", val1)
        fmt.Printf("val1 type: %v\n", reflect.TypeOf(val1))
/*
	val1: 0xc00007c150    <-------------- queuenode.data
	val1 type: reflect.Value
*/



//	func (v Value) Elem() Value
//	Elem returns the value that the interface v contains or that the pointer v points to
        val2 := reflect.ValueOf(&node).Elem().Elem()	// the adress of a Tree element 
        fmt.Printf("val2: %v\n", val2)
        fmt.Printf("val2 type: %v\n", reflect.TypeOf(val2))
/*
	val2: &{1 false []}     <-------------- &Tree
	val2 type: reflect.Value
*/




//	func Indirect(v Value) Value
// 	Indirect returns the value that v points to
	//THIS ALSO GOOD!!! val3 := reflect.Indirect(reflect.ValueOf(&node).Elem().Elem())
	val3 := reflect.ValueOf(&node).Elem().Elem().Elem()
        fmt.Printf("val3: %v\n", val3)
        fmt.Printf("val3 type: %v\n", reflect.TypeOf(val3))
/*
	val3: {1 false []}         <------------- Tree
	val3 type: reflect.Value
*/



	//THIS ALSO GOOD!!!  val4 := reflect.Indirect(reflect.ValueOf(&node).Elem().Elem()).FieldByName("Value")
	val4 := reflect.ValueOf(&node).Elem().Elem().Elem().FieldByName("Value")
	fmt.Printf("val4: %v\n", val4)
	fmt.Printf("val4 type: %v\n", reflect.TypeOf(val4))
/*
	val4: 1                  <--------------- Tree.Value
	val4 type: reflect.Value
*/

}



