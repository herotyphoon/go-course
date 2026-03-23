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

	for i := 0; i < 10; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}

	b := 10

	if b == 10 {
		goto end
	}
end:
	fmt.Println("End")
outerLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				continue outerLoop
			}
			fmt.Println(i, j)
		}
	}
}

