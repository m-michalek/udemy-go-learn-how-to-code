package main

import "fmt"

func main() {
	birthDay := 1993
	for {
		if birthDay > 2021 {
			break
		}

		fmt.Printf("birthDay: %v\n", birthDay)
		birthDay++
	}
}
