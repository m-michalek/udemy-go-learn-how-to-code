package main

import "fmt"

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("custom error: %v\n", ce.info)
}

func foo(e error) {
	fmt.Println(e)
}

func main() {
	err := customErr{info: "irrelevant errors"}

	foo(err)
}
