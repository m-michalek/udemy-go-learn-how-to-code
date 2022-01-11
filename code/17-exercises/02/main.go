package main

import "fmt"

type person struct {
	name string
}

func changeMe(p *person) {
	p.name = "Lenny"
	// (*p).name = "Lenny"
}

func main() {
	p := person{
		name: "Ronald",
	}
	fmt.Printf("p: %v\n", p)

	changeMe(&p)
	fmt.Printf("p: %v\n", p)
}
