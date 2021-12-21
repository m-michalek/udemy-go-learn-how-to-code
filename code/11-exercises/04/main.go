package main

import "fmt"

func main() {
	numbers := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	numbers = append(numbers, 52)
	fmt.Printf("numbers: %v\n", numbers)
	numbers = append(numbers, 53, 54, 55)
	fmt.Printf("numbers: %v\n", numbers)

	y := []int{56, 57, 58, 59, 60}
	numbers = append(numbers, y...)
	fmt.Printf("numbers: %v\n", numbers)
}
