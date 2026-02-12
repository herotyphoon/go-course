package main

import "fmt"

func main() {

	var myString string
	fmt.Println(myString) // This will print an empty string

	myString = "Hello, World!"
	fmt.Println(myString)

	myString = "Hello\nWorld!"
	fmt.Println(myString)

	myString = `Welcome to 

Golang Course!`
	fmt.Println(myString)

	var firstName, lastName string
	firstName = "John"
	lastName = "Doe"

	var fullName string
	fullName = firstName + " " + lastName // Concatenating strings using the + operator
	fmt.Println(fullName)

	fmt.Printf("%s %s\n", firstName, lastName) // Using fmt.Printf to format the output with placeholders

	fullName = fmt.Sprintf("%s %s", firstName, lastName) // Using fmt.Sprintf to create a formatted string
	fmt.Println(fullName)
}