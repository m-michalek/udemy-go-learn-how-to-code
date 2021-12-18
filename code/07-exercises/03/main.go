package main

import "fmt"

const (
	untypedBatchSize       = 12
	typedBatchSize   int64 = 261120
)

func main() {
	fmt.Printf("untypedBatchSize: \t%v\t%T\n", untypedBatchSize, untypedBatchSize)
	fmt.Printf("typedBatchSize: \t%v\t%T\n", typedBatchSize, typedBatchSize)
}
