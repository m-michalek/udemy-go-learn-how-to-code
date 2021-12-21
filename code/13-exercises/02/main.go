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
	p2 := person{
		firstName: "Larissa",
		lastName:  "Miller",
		favouriteIceCream: []string{
			"Peach",
			"Stracialtella",
			"Cookie",
		},
	}

	people := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for k, v := range people {
		fmt.Printf("k: %v\n", k)
		fmt.Printf("v: %v\n", v)
	}
}
