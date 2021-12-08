package main

import "fmt"

type age int

var x age

func main() {
	fmt.Printf("x: %v\n", x)
	fmt.Printf("x: %T\n", x)

	x = 42

	fmt.Printf("x: %v\n", x)
}
