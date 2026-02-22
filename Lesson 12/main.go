package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a uint8 = 205
	var b uint8 = 141

	c := a | b

	fmt.Println("OR Operation Result:")
	fmt.Println(c)
	fmt.Println(strconv.FormatUint(uint64(a), 2))
	fmt.Println(strconv.FormatUint(uint64(b), 2))
	fmt.Println(strconv.FormatUint(uint64(c), 2))

	d := a & b

	fmt.Println("AND Operation Result:")
	fmt.Println(d)
	fmt.Println(strconv.FormatUint(uint64(a), 2))
	fmt.Println(strconv.FormatUint(uint64(b), 2))
	fmt.Println(strconv.FormatUint(uint64(d), 2))

	e := a ^ b

	fmt.Println("XOR Operation Result:")
	fmt.Println(e)
	fmt.Println(strconv.FormatUint(uint64(a), 2))
	fmt.Println(strconv.FormatUint(uint64(b), 2))
	fmt.Println(strconv.FormatUint(uint64(e), 2))

	f := ^a

	fmt.Println("NOT Operation Result:")
	fmt.Println(f)
	fmt.Println(strconv.FormatUint(uint64(a), 2))
	fmt.Println(strconv.FormatUint(uint64(f), 2))

	g := a &^ b

	fmt.Println("AND NOT Operation Result:")
	fmt.Println(g)
	fmt.Println(strconv.FormatUint(uint64(a), 2))
	fmt.Println(strconv.FormatUint(uint64(b), 2))
	fmt.Println(strconv.FormatUint(uint64(g), 2))
}
