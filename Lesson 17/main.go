package main

import "fmt"

func main() {
	name := "Gopher"
	first := name[0]
	fmt.Printf("First byte: %d\n", first)
	last := name[len(name)-1]
	fmt.Printf("Last byte: %d\n", last)
	// inv := name[len(name)]
	// fmt.Printf("Invalid byte: %d\n", inv) // Index out of range error
	// name[1] = 'o' // This will cause a compile-time error: cannot assign to name[1]
	// name = "new"
	// fmt.Printf("Updated name: %s\n", name) // This will work because we are reassigning the entire string, not modifying it in place

	firstLetter := name[0]
	fmt.Println("First letter:", string(firstLetter))

	var b rune = 'G'
	fmt.Println(string(b))

	runeName := []rune(name)
	fmt.Printf("Runes: %v\n", runeName)
	fmt.Printf("First rune: %c\n", runeName[0])

	byteName := []byte(name)
	fmt.Printf("Bytes: %v\n", byteName)
	fmt.Printf("First byte: %d\n", byteName[0])

	fmt.Println(name[:5])
	fmt.Println(name[1:4])
	fmt.Println(name[1:])
	fmt.Println(name[:])

	newName := "コーディング"
	fmt.Println(newName)
	firstChar := newName[0]
	fmt.Println(string(firstChar))

	unicodeString := []rune(newName)
	fmt.Printf("Unicode runes: %v\n", unicodeString)
	fmt.Printf("First character: %c\n", unicodeString[0])
}
