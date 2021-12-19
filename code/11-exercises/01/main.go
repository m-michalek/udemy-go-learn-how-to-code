package main

import "fmt"

func main() {
	numbers := [5]int{10, 20, 30, 40, 50}

	for _, number := range numbers {
		fmt.Printf("number: %v\n", number)
	}

	fmt.Printf("numbers: %T\n", numbers)

}
