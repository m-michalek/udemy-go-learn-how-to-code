package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	m := sync.Mutex{}

	interations := 100
	incrementer := 0
	wg.Add(interations)

	fmt.Printf("incrementer: %v\n", incrementer)
	fmt.Printf("NumCPU: %v\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine: %v\n", runtime.NumGoroutine())

	for i := 0; i < interations; i++ {
		go func() {
			// m.Lock()
			// v := incrementer
			// v++
			// incrementer = v
			// fmt.Printf("new incrementer: %v\n", incrementer)
			// m.Unlock()
			// wg.Done()

			m.Lock()
			incrementer++
			fmt.Printf("new incrementer: %v\n", incrementer)
			m.Unlock()

			fmt.Printf("current NumGoroutine: %v\n", runtime.NumGoroutine())
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("---------------------------")
	fmt.Printf("end value of incrementer: %v\n", incrementer)

	// in order to run the program with race detection execute the following commands in the directory 04:
	// go run -race main.go

	// something like "Found 1 data race(s)" should be missing at the end of the output
	// so there is no data race
}
