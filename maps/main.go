package main

import "fmt"

func main() {

    m := make(map[string]int)

    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map: ", m)

    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    fmt.Println("len: ", len(m))


    v2, prs := m["k2"]
    fmt.Println("prs before: ", prs)
    fmt.Println("v2 before: ", v2)
    delete(m, "k2")
    fmt.Println("map: ", m)

    _, prs = m["k2"]
    fmt.Println("prs: ", prs)



    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map: ", n)

}
