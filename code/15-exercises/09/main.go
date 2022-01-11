package main

import "fmt"

func sum(f func(i []int) int, i []int) int {
	return f(i)
}

func main() {
	total := sum(func(i []int) int {
		sum := 0
		for _, x := range i {
			sum += x
		}
		return sum
	}, []int{1, 2, 3})

	fmt.Printf("total: %v\n", total)
}
