package word

import (
	"fmt"
	"testing"

	"example.com/exercise/02/quote"
)

func TestCount(t *testing.T) {
	n := Count(quote.SunAlso)
	if n != 1349 {
		t.Errorf("Expected: %v, Got: %v", n, 1349)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		Count(quote.SunAlso)
	}
}

func ExampleCount() {
	fmt.Println(Count("hello world"))

	// Output:
	// 2
}

func TestUseCount(t *testing.T) {
	m := UseCount(quote.SunAlso)
	switch {
	case m["one"] != 8:
		t.Errorf("Expected: %v, Got: %v", 8, m["one"])
	case m["two"] != 6:
		t.Errorf("Expected: %v, Got: %v", 6, m["two"])
	case m["three"] != 5:
		t.Errorf("Expected: %v, Got: %v", 5, m["three"])
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func ExampleUseCount() {
	fmt.Println(UseCount("hello world"))

	// Output:
	// map[hello:1 world:1]
}
