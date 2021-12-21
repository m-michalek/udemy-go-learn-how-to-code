package main

import "fmt"

type person struct {
	firstName         string
	lastName          string
	favouriteIceCream []string
}

func main() {
	p1 := person{
		firstName: "Jon",
		lastName:  "Doe",
		favouriteIceCream: []string{
			"Stawberry",
			"Chocolate",
			"Vanilla",
		},
	}
	fmt.Printf("p1: %v\n", p1)
	for _, ice := range p1.favouriteIceCream {
		fmt.Printf("ice: %v\n", ice)
	}

	p2 := person{
		firstName: "Larissa",
		lastName:  "Doe",
		favouriteIceCream: []string{
			"Peach",
			"Stracialtella",
			"Cookie",
		},
	}
	fmt.Printf("p2: %v\n", p2)
	for _, ice := range p2.favouriteIceCream {
		fmt.Printf("ice: %v\n", ice)
	}
}
