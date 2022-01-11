package main

import (
	"fmt"
	"math/rand"
	"time"
)

func add() func() []int {
	list := []int{}
	return func() []int {
		x := rand.New(rand.NewSource(time.Now().UnixNano()))

		list = append(list, x.Intn(100))
		return list
	}
}

func main() {
	x := add()
	fmt.Printf("x: %v\n", x())
	fmt.Printf("x: %v\n", x())
	fmt.Printf("x: %v\n", x())

	y := add()
	fmt.Printf("y: %v\n", y())
	fmt.Printf("y: %v\n", y())
	fmt.Printf("y: %v\n", y())
}
