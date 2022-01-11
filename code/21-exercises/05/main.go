package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	wg := sync.WaitGroup{}

	interations := 100
	var incrementor int64
	wg.Add(interations)

	fmt.Printf("incrementor: %v\n", incrementor)
	fmt.Printf("NumCPU: %v\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine: %v\n", runtime.NumGoroutine())

	for i := 0; i < interations; i++ {
		go func() {
			atomic.AddInt64(&incrementor, 1)
			fmt.Printf("new incrementor: %v\n", atomic.LoadInt64(&incrementor))
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("---------------------------")
	fmt.Printf("end value of incrementor: %v\n", incrementor)

	// in order to run the program with race detection execute the following commands in the directory 04:
	// go run -race main.go

	// something like "Found 1 data race(s)" should be missing at the end of the output
	// so there is no data race
}
