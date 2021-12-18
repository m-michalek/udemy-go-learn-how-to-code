package main

import "fmt"

// is a statement
// 0 == 0

// is an expression
// a := 0 == 0

func main() {
	a := "" == ``
	fmt.Printf("a: %v\n", a)

	b := 0 <= 1
	fmt.Printf("b: %v\n", b)

	c := 1 >= 0
	fmt.Printf("c: %v\n", c)

	d := "a" != "b"
	fmt.Printf("d: %v\n", d)

	e := 0 < 1
	fmt.Printf("e: %v\n", e)

	f := 1 > 0
	fmt.Printf("f: %v\n", f)
}
