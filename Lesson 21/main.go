package main

import "fmt"

// package level variable.
var a = 10

func main() {
	// function level variable.
	b := 20

	fmt.Println(a)
	fmt.Println(b)

	{
		// block variable.
		c := 30
		// a, b & c can be accessed here.

		fmt.Println(a)
		fmt.Println(c)
		fmt.Println(c)
	}
	// a & b can be accessed here
	// but c cannot be accessed here

	x := 20

	if true {
		x := 10        // This will shadow the outer 'x'.
		fmt.Println(x) // This will print 10 not 20
	}

	fmt.Println(x) // This will print 20, the outer 'x'.
}
