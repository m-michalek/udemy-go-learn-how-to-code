package main

import "fmt"

func main() {
	x := func() func() int {
		return func() int {
			return 42
		}
	}()

	fmt.Printf("x: %v\n", x())
}
