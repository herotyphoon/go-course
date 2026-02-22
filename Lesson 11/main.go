package main

import "fmt"

func main() {
	a := 20
	b := 5
	c := 15.50

	sum := a + b
	difference := a - b
	product := a * b
	quotient := a / b
	remainder := a % b

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)
	fmt.Println("Remainder:", remainder)

	firstName := "John"
	lastName := "Doe"

	fullName := firstName + " " + lastName // Also called string concatenation

	fmt.Println("Full Name:", fullName)
	fmt.Printf("%s %s %s", firstName, lastName, "\n")

	complex1 := complex(10, 15) // Also can be written as 10 + 15i
	complex2 := complex(5, 10)  // Also can be written as 5 + 10i

	complexSum := complex1 + complex2
	complexDifference := complex1 - complex2
	complexProduct := complex1 * complex2
	complexQuotient := complex1 / complex2

	fmt.Println("Complex Sum:", complexSum)
	fmt.Println("Complex Difference:", complexDifference)
	fmt.Println("Complex Product:", complexProduct)
	fmt.Println("Complex Quotient:", complexQuotient)

	newSum := sum + int(c) // Type conversion from float64 to int
	fmt.Println("New Sum (with type conversion):", newSum)

	a += 10 // a = a + 10
	fmt.Println("Updated a:", a)

	a -= 5 // a = a - 5
	fmt.Println("Updated a:", a)

	a *= 2 // a = a * 2
	fmt.Println("Updated a:", a)

	a /= 4 // a = a / 4
	fmt.Println("Updated a:", a)

	a %= 3 // a = a % 3
	fmt.Println("Updated a:", a)

	// zeroDivision := 10 / 0 // This will cause a runtime panic: division by zero
	// fmt.Println("Zero Division Result:", zeroDivision)

	a++ // a = a + 1
	fmt.Println("After Incrementing a:", a)

	a-- // a = a - 1
	fmt.Println("After Decrementing a:", a)
}
