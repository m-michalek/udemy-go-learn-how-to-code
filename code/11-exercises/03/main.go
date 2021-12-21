package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Printf("numbers: %v\n", numbers[:5])
	fmt.Printf("numbers: %v\n", numbers[5:])
	fmt.Printf("numbers: %v\n", numbers[2:7])
	fmt.Printf("numbers: %v\n", numbers[1:6])
}
