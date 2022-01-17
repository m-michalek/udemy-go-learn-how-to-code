package main

import (
	"fmt"

	"example.com/exercise/01/dog"
)

func main() {
	dogYears := dog.Years(10)

	fmt.Printf("dogYears: %v\n", dogYears)

}
