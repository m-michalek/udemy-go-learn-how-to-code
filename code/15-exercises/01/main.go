package main

import "fmt"

func main() {
	n1 := foo()
	fmt.Printf("n1: %v\t%T\n", n1, n1)

	n2, s1 := bar()
	fmt.Printf("n2: %v\t%T\n", n2, n2)
	fmt.Printf("s1: %v\t%T\n", s1, s1)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 42, "42"
}
