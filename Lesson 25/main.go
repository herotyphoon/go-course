package main

import "fmt"

func main() {
	sayHi()
	sayHiToSomeone("Typhoon")
	fmt.Println(fullName("Hero", "Typhoon"))
	fmt.Println(fullNameWithLength("Hero", "Typhoon"))
	_, l := fullNameWithLength("Bruce", "Wayne") // _ is ignored
	fmt.Println(l)
}

func sayHi() {
	fmt.Println("Hi")
}

func sayHiToSomeone(name string) {
	fmt.Println("Hi", name)
}

func fullName(firstName, lastName string) string {
	return fmt.Sprintf("%s %s", firstName, lastName)
}

func fullNameWithLength(firstName, lastName string) (string, int) {
	fullName := fmt.Sprintf("%s %s", firstName, lastName)
	return fullName, len(fullName)
}
