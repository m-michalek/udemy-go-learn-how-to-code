package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int)

	go func(c1 chan<- int) {
		c1 <- 42
	}(c1)
	fmt.Println(<-c1)

	fmt.Printf("------\n")
	fmt.Printf("c1\t%T\n", c1)

	c2 := make(chan int)

	go func(c2 chan<- int) {
		c2 <- 42
	}(c2)
	fmt.Println(<-c2)

	fmt.Printf("------\n")
	fmt.Printf("c2\t%T\n", c2)

}
