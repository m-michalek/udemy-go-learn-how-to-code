package main

import (
	"fmt"

	"example.com/exercise/02/quote"
	"example.com/exercise/02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
