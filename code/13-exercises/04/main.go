package main

import "fmt"

func main() {
	t := struct {
		doors     int
		color     string
		fourWheel bool
		manual    map[string]string
		owners    []string
	}{
		doors:     4,
		color:     "red",
		fourWheel: true,
		manual: map[string]string{
			"driving": "In order to drive...",
			"parking": "In order to park...",
			"honking": "In order to honk...",
		},
		owners: []string{
			"Jon Doe",
			"Mike Miller",
			"Henry Ford",
		},
	}

	fmt.Printf("t: %#v\n", t)
	fmt.Println(t.fourWheel)
}
