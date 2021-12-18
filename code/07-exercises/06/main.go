package main

import "fmt"

const (
	year01 = 2018 + iota
	year02
	year03
	year04
)

func main() {
	fmt.Printf("year01: %v\n", year01)
	fmt.Printf("year02: %v\n", year02)
	fmt.Printf("year03: %v\n", year03)
	fmt.Printf("year04: %v\n", year04)

}
