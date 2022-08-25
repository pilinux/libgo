// Package datamath provides mathematical functions which are not implemented in the standard library.
package datamath_test

import (
	"fmt"
	"testing"

	"github.com/pilinux/libgo/datamath"
)

type fibonacciNthTermTest struct {
	input, expected int
}

var fibonacciNthTermTests = []fibonacciNthTermTest{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{8, 21},
	{25, 75025},

	{-0, 0},
	{-1, -1},
	{-2, -1},
	{-3, -2},
	{-8, -21},
	{-25, -75025},
}

func TestFibonacciNthTerm(t *testing.T) {
	for _, test := range fibonacciNthTermTests {
		result := datamath.FibonacciNthTerm(test.input)
		expected := test.expected
		if result != expected {
			t.Errorf(
				"got: %v\nwant: %v",
				result, expected,
			)
		}
	}
}

func ExampleFibonacciNthTerm() {
	fmt.Println(datamath.FibonacciNthTerm(25))
	// Output: 75025
}
