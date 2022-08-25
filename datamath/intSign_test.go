package datamath_test

import (
	"fmt"
	"testing"

	"github.com/pilinux/libgo/datamath"
)

type intSignTest struct {
	input    int
	expected bool
}

var intSignTests = []intSignTest{
	{0, false},
	{-0, false},
	{1, false},
	{-1, true},
	{9223372036854775807, false}, // max value of 64-bit int
	{-9223372036854775808, true}, // min value of 64-bit int
}

func TestIntSign(t *testing.T) {
	for _, test := range intSignTests {
		result := datamath.IntSign(test.input)
		expected := test.expected
		if result != expected {
			t.Errorf(
				"got: %v\nwant: %v",
				result, expected,
			)
		}
	}
}

func ExampleIntSign() {
	fmt.Println(datamath.IntSign(-1200))
	// Output: true
}
