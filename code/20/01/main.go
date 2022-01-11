package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func init() {
	fmt.Println("I'm running without explicit invocation before the main function runs 👌")
}

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1)

	go foo()
	bar()

	wg.Wait()
	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

}

func foo() {
	for i := 1; i <= 100; i++ {
		// fmt.Println(i)
	}
	wg.Done()

	fmt.Println("foo done")
}

func bar() {
	for i := 1; i <= 100; i++ {
		// fmt.Println(i)
	}
	fmt.Println("bar done")
}
