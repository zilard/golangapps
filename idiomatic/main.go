package main

// constant declarations
// variable grouping
// functions that panic
// struct initialization
// mutex grouping
// interface declarations/naming
// function grouping
// http handler naming
// enums
// consturctor

// in golang we treat constants just like other variables
// no uppercase for constants
//const scalar = 0.1

// if you want to export it then we can capitalize
// the first letter
// const Scalar = 0.1

// if you wanna have multiple constants
// then you need to group them, on top of the file
/*
const (
	Scalar  = 0.1
	Version = 0.1
)
*/

// variable grouping
// if I have multiple variables in a function
// then I'm always grouping them in a var
// because it's much more idiomatic
/*
func Foo() int {
	var (
		x   = 100
		y   = 2
		foo = "foo"
	)

	fmt.Println(foo)

	return x + y
}
*/

// functions that panic
// if your function panics,
// then you always prefix it with a "Must"
func MustParseIntFromString(s string) int {
	// logic
	panic("oops")

	return 10
}

// struct initialization

func main() {

}
