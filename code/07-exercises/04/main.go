package main

import "fmt"

func main() {
	number := 42
	fmt.Printf("number: \t%d\t%b\t%#x\t\n", number, number, number)

	shiftedNumber := number << 1
	fmt.Printf("shiftedNumber: \t%d\t%b\t%#x\t\n", shiftedNumber, shiftedNumber, shiftedNumber)

}
