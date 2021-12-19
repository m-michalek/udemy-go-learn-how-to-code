package main

import "fmt"

func main() {
	name := "James Bond"

	if name == "Moneypenny" {
		fmt.Printf("name: %v\n", name)
	} else if name == "James Bond" {
		fmt.Println("You are", name)
	} else {
		fmt.Println("I dont know you!")
	}
}
