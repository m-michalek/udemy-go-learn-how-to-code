package main

import "fmt"

func main() {
	c := make(chan int)

	go send(c)

	receive(c)
}

func send(cs chan<- int) {
	fmt.Printf("cs: %v\t%T\n", cs, cs)
	for i := 0; i < 100; i++ {
		cs <- i
	}
	close(cs)
}

func receive(cr <-chan int) {
	for v := range cr {
		fmt.Printf("v: %v\n", v)
	}
}
