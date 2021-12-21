package main

import "fmt"

func main() {
	x := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Hellooooo, James"},
	}

	for _, record := range x {
		fmt.Printf("record: %v\n", record)
		for _, data := range record {
			fmt.Println(data)
		}
	}
}
