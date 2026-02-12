package main

import "fmt"

func main() {

	// Variable of type uint8
	var smallPositiveValue uint8
	smallPositiveValue = 10
	fmt.Println(smallPositiveValue)

	// Variable of type int8
	var smallNegativeInt int8
	smallNegativeInt = -10
	fmt.Println(smallNegativeInt)

	// uint16, uint32, uint64
	// int16, int32, in64
	var myInt int = 230000043
	fmt.Println(myInt)

	myInt = int(smallNegativeInt)
	myInt = int(smallPositiveValue)

	// to typecase a value int()

	// byte is an alias for uint8, and it is mainly used to represent raw data
	//and also to distinguish it from other uint8 values that may represent something else, such as a small positive integer.
	var myByte byte = 'A'
	fmt.Println(myByte)

	// rune is an alias for int32, and it is used to represent Unicode code points. It can represent any character in the Unicode standard, including letters, digits, symbols, and emojis.
	var myRune rune = 'B'
	myRune = '😊'
	fmt.Println(myRune)
	
}
