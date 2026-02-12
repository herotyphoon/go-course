package main

import "fmt"

func main() {

	var myString string

	// Strings in Go are initialized to the empty string ("") by default
	fmt.Println(myString)

	myString = "Hello, World!"
	fmt.Println(myString)

	var myInt int

	// Integers in Go are initialized to 0 by default
	fmt.Println(myInt)

	myInt = 10
	fmt.Println(myInt)

	var myBool bool

	// Booleans in Go are initialized to false by default
	fmt.Println(myBool)

	myBool = true
	fmt.Println(myBool)
}
