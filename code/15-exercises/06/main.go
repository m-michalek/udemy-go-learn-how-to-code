package main

import "fmt"

func main() {
	func() { fmt.Println("build and use an anonymous func...") }()
}
