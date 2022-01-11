package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	incrementer := 0
	fmt.Printf("incrementer: %v\n", incrementer)
	fmt.Printf("NumCPU: %v\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine: %v\n", runtime.NumGoroutine())

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			fmt.Printf("new incrementer: %v\n", incrementer)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("---------------------------")
	fmt.Printf("end value of incrementer: %v\n", incrementer)

	// in order to run the program with race detection execute the following commands in the directory 03:
	// go run -race main.go

	// lookout for something like "Found 1 data race(s)" at the end of the output
}
