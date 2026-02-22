package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a uint8 = 20
	result := a >> 2

	fmt.Println("Right Shift Result:")
	fmt.Println(result)
	fmt.Println("Binary Representation:")
	fmt.Println(strconv.FormatUint(uint64(a), 2))
	fmt.Println(strconv.FormatUint(uint64(result), 2))

	result = a << 2

	fmt.Println("Left Shift Result:")
	fmt.Println(result)
	fmt.Println("Binary Representation:")
	fmt.Println(strconv.FormatUint(uint64(a), 2))
	fmt.Println(strconv.FormatUint(uint64(result), 2))

	// Bit Mask
	a = 6
	a <<= 1
	fmt.Println("Bit Mask Result:")
	fmt.Println(a)
	fmt.Println("Binary Representation:")
	fmt.Println(strconv.FormatUint(uint64(a), 2))

	// Set Bit at a specific position
	a = 0
	position := 3
	a |= (1 << position)
	fmt.Println("Set Bit Result:")
	fmt.Println(a)
	fmt.Println("Binary Representation:")
	fmt.Println(strconv.FormatUint(uint64(a), 2))

	// Clear Bit at a specific position
	a = 15
	position = 2
	a &^= (1 << position)
	fmt.Println("Clear Bit Result:")
	fmt.Println(a)
	fmt.Println("Binary Representation:")
	fmt.Println(strconv.FormatUint(uint64(a), 2))

	// Toggle Bit at a specific position
	a = 5
	position = 1
	a ^= (1 << position)
	fmt.Println("Toggle Bit Result:")
	fmt.Println(a)
	fmt.Println("Binary Representation:")
	fmt.Println(strconv.FormatUint(uint64(a), 2))
}
