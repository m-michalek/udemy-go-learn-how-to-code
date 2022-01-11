package main

import (
	"fmt"
)

func main() {
	c1 := make(chan int)

	go func() {
		c1 <- 42
	}()

	fmt.Println(<-c1)

	c2 := make(chan string, 2)

	c2 <- "buffered"
	c2 <- "channel"

	fmt.Println(<-c2)
	fmt.Println(<-c2)
}
