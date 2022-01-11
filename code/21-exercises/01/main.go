package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go foo()
	go bar()

	fmt.Println("NumGoroutine:", runtime.NumGoroutine())

	wg.Wait()
}

func foo() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("foo: %v\n", i)
	}
	wg.Done()
}

func bar() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("bar: %v\n", i)
	}
	wg.Done()
}
