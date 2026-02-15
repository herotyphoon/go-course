package main

import "fmt"

// var myInt int = 10

func main() {
	// Group Similar Variables
	// var (
	// 	myInt    = 10
	// 	myString = "Code"
	// )
	// Explicit Type Assignment
	// var myInt int = 10
	// var myString string = "Code"

	// Implicit Type Assignment
	// var myInt = 10
	// var myString = "Code"

	// Multiple Variable Declaration
	// var month, year = 2, 12

	// Shorthand Declaration
	age := 25

	// age:= 30 Shows error

	age, year := 30, 2026

	fmt.Println(age)
	fmt.Println(year)

	year = 2025
	fmt.Println(year)

	// fmt.Println(myInt)
	// fmt.Println(myString)
	// fmt.Println(month)
	// fmt.Println(year)
	// something()

	// All Examples
	// var age int
	// var age int = 10
	// var age = 10
	// var age, name = 10, "Code"
	// age := 25
	// age, name := 10, "Code"
}

// func something() {
// 	fmt.Println(myInt)
// }
