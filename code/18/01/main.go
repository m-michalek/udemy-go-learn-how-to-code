package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {
	p1 := person{
		Name: "Ron",
		Age:  22,
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		panic("marshalling person went wrong")
	}

	fmt.Printf("marshalled person: %+v\t%T\t%v\n", bs, bs, string(bs))

	var p2 person
	err = json.Unmarshal(bs, &p2)
	if err != nil {
		panic("unmarshalling person went wrong")
	}

	fmt.Printf("unmarshalled person: %+v\t%T\n", p2, p2)
}
