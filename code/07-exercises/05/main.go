package main

import "fmt"

func main() {
	var rawStringLiteral = `"this is printed in quotes and all escaped character are shown \t \n"`
	fmt.Printf("rawStringLiteral: %v\n", rawStringLiteral)
}
