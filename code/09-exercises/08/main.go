package main

import "fmt"

func main() {
	switch {

	case false:
		fmt.Println("This is not getting printed")
	case true:
		fmt.Println("This gets printed")
	}
}
