package main

import (
	"fmt"

	"example.com/exercise/01/dog"
)

func main() {
	rex := dog.Dog{
		Name:   "rex",
		DogAge: dog.Years(2),
	}

	fmt.Printf("rex: %+v\n", rex)
}
