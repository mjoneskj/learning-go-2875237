package main

import (
	"fmt"
)

func main() {
	theAnswer := 42
	var result string
	if theAnswer < 0 {
		result = "Less than zero"

	} else if theAnswer == 0 {
		result = "Equal to zero"
	} else {
		result = "Great than zero"
	}
	fmt.Println(result)
	fmt.Println("Conditional logic")

	if x := -42; x < 0 {
		result = "Less than zero"
	} else if x == 0 {
		result = "Less equals zero"
	} else {
		result = "Greater than zero"
	}
	fmt.Println(result)

}
