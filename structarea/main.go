package main

import "fmt"

type rect struct {
    width, height int
}

// methods can be defined for eithr pointr or value receiver

// this method is using "pointer receiver type"   (r *rect)
func (r *rect) area() int {
    return r.width * r.height
}

// Go automatically handles conversion between values and pointers for method calls
// you may want to use a pointer receivr type to avoid copying on method calls
// or to allow the method to mutate the receiving struct


// this method is using "value receiver type"   (r rect)
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}

    fmt.Println("area: ", r.area())
    fmt.Println("perim: ", r.perim())

    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim: ", rp.perim())

}


