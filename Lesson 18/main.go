package main

import "fmt"

func main() {
	nameAge := map[string]int{
		"Foo": 25,
		"Bar": 20,
	}

	fmt.Println(nameAge)

	firstNameLastName := map[string]string{
		"Foo": "Bar",
		"Bar": "Foo",
	}
	fmt.Println(firstNameLastName)

	var nameAge2 map[string]int

	fmt.Println(nameAge2)

	nameAge3 := map[string]int{}
	nameAge4 := make(map[string]int)
	var nameAge5 map[string]int = map[string]int{}

	fmt.Println(nameAge3)
	fmt.Println(nameAge4)
	fmt.Println(nameAge5)

	fooAge, ok := nameAge["foo"]
	if !ok {
		fmt.Println("not found")
	} else {
		fmt.Println(fooAge)
	}

	fooAge1, ok := nameAge["Foo"]
	if !ok {
		fmt.Println("not found")
	} else {
		fmt.Println(fooAge1)
	}

	nameAge["Foo"] = 20

	fmt.Println(nameAge)

}
