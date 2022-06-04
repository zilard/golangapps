package main


import (
	"fmt"
)


type intf1 interface {
	afun()
}


type intf2 interface {
	bfun()
}


type intf3 interface {
	cfun()
}


type struct1 struct {
	data string
}


func (s struct1) afun() string {
	return fmt.Sprintf("afun: %v\n", s.data)
}

func (s struct1) bfun() string {
        return fmt.Sprintf("bfun: %v\n", s.data)
}

func (s struct1) cfun() string {
        return fmt.Sprintf("cfun: %v\n", s.data)
}


func main() {


	sl := make([]int, 2, 5)

	//sl := []int{0,0,0}

	fmt.Printf("slice: %v\n", sl)



/*

	var a [4]int                      // array with zero values
	var b [47]int = [4]int{0, 1, 2}    // partially initialized array
	var c [4]int = [4]int{1, 2, 3, 4} // array initialization
	d := [...]int{5, 6, 7, 0}         // ... - means that array size equals the number of elements in the array literal

	fmt.Printf("a: length: %d, capacity: %d, data: %v\n", len(a), cap(a), a)
	fmt.Printf("b: length: %d, capacity: %d, data: %v\n", len(b), cap(b), b)
	fmt.Printf("c: length: %d, capacity: %d, data: %v\n", len(c), cap(c), c)
	fmt.Printf("d: length: %d, capacity: %d, data: %v\n", len(d), cap(d), d)
*/

	s1 := struct1{ data: "hello" }

	fmt.Printf("s1 calling cfun: %v\n", s1.cfun())

	var intf []interface{}

        intf = make([]interface{}, 3)

	intf = append(intf, s1)


	fmt.Printf("intf %v\n", intf)


}


