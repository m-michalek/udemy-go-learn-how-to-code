package main

import "fmt"

type person struct{}

func (p *person) speak() {
	fmt.Println("hello")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{}

	saySomething(&p1)
}
