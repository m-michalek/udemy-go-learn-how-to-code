package mymath

import (
	"fmt"
	"testing"
)

var centeredAvgTestData = []struct {
	xi       []int
	expected float64
}{
	{xi: []int{1, 4, 6, 8, 100}, expected: 6},
	{xi: []int{0, 8, 10, 1000}, expected: 9},
	{xi: []int{9000, 4, 10, 8, 6, 12}, expected: 9},
	{xi: []int{123, 744, 140, 200}, expected: 170},
}

func TestCenteredAvg(t *testing.T) {
	for _, v := range centeredAvgTestData {
		ca := CenteredAvg(v.xi)
		if ca != v.expected {
			t.Errorf("Expected: %v, Got: %v", ca, v.expected)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 4, 6, 8, 100}))

	// Output:
	// 6
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 4, 6, 8, 100})
	}
}
