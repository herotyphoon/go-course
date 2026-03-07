package main

import "fmt"

func isEven(num int) bool {
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	age := 10
	if age >= 18 {
		fmt.Println("You are an adult!")
	}

	if age >= 18 {
		fmt.Println("You are an adult!")
	} else {
		fmt.Println("You are not an adult!")
	}

	if age >= 18 {
		fmt.Println("You are an adult!")
	} else if age >= 13 {
		fmt.Println("You are a teenager!")
	} else {
		fmt.Println("You are a child!")
	}

	if even := isEven(age); even {
		fmt.Println("Age is even")
	}

}
