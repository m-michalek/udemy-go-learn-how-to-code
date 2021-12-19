package main

import "fmt"

func main() {
	birthDay := 1993
	for birthDay <= 2021 {
		fmt.Printf("birthDay: %v\n", birthDay)

		birthDay++
	}
}
