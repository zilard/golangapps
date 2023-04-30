package main

import "fmt"

func myExecuteFunc(s string) {
	fmt.Println("my ex func", s)
}

func main() {
	Execute(myExecuteFunc)
}

// type function
// this is coming from a third party lib
type ExecuteFn func(string)

func Execute(fn ExecuteFn) {
	fn("FOO BAR BAZ")

}
