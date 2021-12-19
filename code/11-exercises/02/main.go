package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	for _, number := range numbers {
		fmt.Printf("number: %v\n", number)
	}

	fmt.Printf("numbers: %T\n", numbers)

}
