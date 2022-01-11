package main

import (
	"fmt"
)

func main() {
	cr := gen()
	receive(cr)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive(cr <-chan int) {
	for v := range cr {
		fmt.Println(v)
	}
}
