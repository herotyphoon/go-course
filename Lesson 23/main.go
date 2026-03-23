package main

import "fmt"

func main() {
	a := 100

	for {
		fmt.Println(a)
		a--
		if a == 90 {
			break
		}
	}
}

