package main

import (
	kata "MultiplesOf3or5"
	"fmt"
)

func main() {
	fmt.Println("Enter a number and press enter or press ctrl+c to exit: ")
	for {
		var number int
		fmt.Scanln(&number)
		fmt.Printf("%d\n", kata.Multiple3And5(number))
	}
}
