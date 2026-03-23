package main

import "fmt"

func main() {
	a := 10

	for a > 0 {
		fmt.Println(a)
		a--
	}

	for {
		fmt.Println(a)
		a--
		if a < -1000 {
			break
		}
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	b := []int{1, 2, 3, 4, 5}

	for index, value := range b {
		fmt.Printf("index : %d, value: %d\n", index, value)
	}

	ages := map[string]int{
		"Alice": 20,
		"Bob":   25,
	}

	for name, age := range ages {
		fmt.Printf("%s is %d years old.\n", name, age)
	}
}
