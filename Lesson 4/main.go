package main

import "fmt"

func main() {
	var smallFloat float32
	fmt.Println(smallFloat)

	smallFloat = 23.0234324
	fmt.Println(smallFloat)

	var bigFloat float64
	fmt.Println(bigFloat)

	bigFloat = 43.08094938590435
	fmt.Println(bigFloat)

	bigFloat = 43.08094938590435567863456433
	fmt.Println(bigFloat)

	var myComplex complex128
	fmt.Println(myComplex)
	myComplex = complex(bigFloat, bigFloat)
	fmt.Println(myComplex)

	var myRealPart, myImaginaryPart float64
	myRealPart = real(myComplex)
	myImaginaryPart = imag(myComplex)
	fmt.Println("Real part:", myRealPart)
	fmt.Println("Imaginary part:", myImaginaryPart)
}
