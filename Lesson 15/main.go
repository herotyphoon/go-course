package main

import "fmt"

func main() {
	var a []int
	fmt.Println(a == nil)

	var b = []int{}
	fmt.Println(b == nil)

	c := []int{1, 2, 3, 4, 5}
	fmt.Println(c)

	d := []int{10, 2: 15, 20}
	fmt.Println(d)

	e := make([]int, 5)
	fmt.Println(e)

	fmt.Println(len(c))
	fmt.Println(cap(c))

	f := make([]string, 5, 10)
	fmt.Println(f)
	fmt.Println(len(f))
	fmt.Println(cap(f))

	e = append(e, 10)
	fmt.Println(e)

	e = append(e, 5)
	fmt.Println(e)

	e = append(e, c...)
	fmt.Println(e)
}
