package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	c := [5]int{10, 2: 20, 50}
	fmt.Println(c)

	d := [...]int{1, 2, 3, 4, 5}
	fmt.Println(d)

	fmt.Println(len(d))

	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])
	fmt.Println(a[3])
	fmt.Println(a[4])

	a[0] = 10
	a[1] = 20
	a[2] = 30
	a[3] = 40
	a[4] = 50

	fmt.Println(a)

	e := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(e)
}
