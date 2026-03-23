package main

import "fmt"

func main() {
	dayNum := 1

	switch dayNum {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	word := "Wow"

	switch wordLen := len(word); wordLen {
	case 2:
		fmt.Println("Word Length is 2")
	case 3:
		fmt.Println("Word Length is 3")
	default:
		fmt.Printf("Word Length is %s", wordLen)
	}

	word2 := "a"

	switch word2 {
	case "a", "b":
		fmt.Println("a or b")
	case "c", "d":
		fmt.Println("c or d")
	default:
		fmt.Println("Not a, b, c or d")
	}

	switch word2 {
	case "a":
		fallthrough
	case "b":
		fmt.Println("a or b")
	}
}
