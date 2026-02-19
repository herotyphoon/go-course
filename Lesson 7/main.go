package main

import "fmt"

// const x = 10
// const y int32 = 15

const (
	x               = 10
	y         int32 = 15
	appName         = "Wow"
	isRunning       = false
	charc           = 'a'
	isTrue          = 1 < 2
)

func main() {
	var a int
	a = x
	fmt.Println(a)

	var b float64
	b = x
	fmt.Println(b)

	var c int
	c = int(y) // Typed const requires explicit type casting for assigning
	fmt.Println(c)

	var d int32
	d = y // Same types so type casting not required
	fmt.Println(d)

	// const z = a + c Does not work as value on RHS are variables

	// These are allowed
	const z = complex(100, 101.5)
	const l = imag(z)
}
