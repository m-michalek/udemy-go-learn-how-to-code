package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("My name is", p.first, p.last, "and I am", p.age, "years old")
}

func main() {
	p1 := person{first: "Jon", last: "Doe", age: 22}

	p1.speak()
}
