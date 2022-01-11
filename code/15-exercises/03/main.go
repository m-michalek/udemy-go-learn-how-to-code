package main

import "fmt"

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}

func third() {
	fmt.Println("third")
}

func main() {
	first()
	defer second()
	third()
}
