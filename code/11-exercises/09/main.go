package main

import "fmt"

func main() {
	x := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "...", "Cars"},
		"moneypenny_miss": {"Helloooo, Bond", "..."},
		"doe_jon":         {"Some stuff", "..."},
	}

	x["sanchez_rick"] = []string{"sience", "beer"}

	for k, v := range x {
		fmt.Printf("k: %v\n", k)
		for i, item := range v {
			fmt.Println(i, item)
		}
		fmt.Println()
	}
}
