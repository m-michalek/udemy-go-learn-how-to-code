package main

import "fmt"

func main() {
	numbers := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	numbers = append(numbers[:3], numbers[6:]...)
	fmt.Printf("numbers: %v\n", numbers)
}
