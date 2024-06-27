package main

import (
	kata "ValidBraces"
	"fmt"
)

func main() {
	fmt.Println("Enter input or press ctrl+c to exit: ")
	for {
		fmt.Println(">")
		var input string
		fmt.Scanln(&input)
		if kata.ValidBraces(input) {
			fmt.Println("String is valid")
		} else {
			fmt.Println("String is not valid")
		}
	}
}
