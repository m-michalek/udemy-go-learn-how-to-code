package main

import "fmt"

func main() {
	sum1 := foo([]int{1, 2, 3, 4, 5}...)
	fmt.Printf("sum: %v\n", sum1)

	sum2 := bar([]int{1, 2, 3, 4, 5})
	fmt.Printf("sum: %v\n", sum2)
}

func foo(x ...int) int {
	var sum int

	for _, i := range x {
		sum += i
	}

	return sum
}

func bar(x []int) int {
	var sum int

	for _, i := range x {
		sum += i
	}

	return sum
}
