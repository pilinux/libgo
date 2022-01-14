// Package datamath provides mathematical functions which are not implemented in the standard library.
package datamath_test

import (
	"testing"

	"github.com/pilinux/libgo/datamath"
)

func TestFibonacciNthTerm(t *testing.T) {
	term0 := 0
	term1 := 1
	term2 := 2
	term3 := 3
	term8 := 8
	term25 := 25

	termN0 := -0
	termN1 := -1
	termN2 := -2
	termN3 := -3
	termN8 := -8
	termN25 := -25

	expected0 := 0
	expected1 := 1
	expected2 := 1
	expected3 := 2
	expected8 := 21
	expected25 := 75025

	expectedN0 := 0
	expectedN1 := -1
	expectedN2 := -1
	expectedN3 := -2
	expectedN8 := -21
	expectedN25 := -75025

	result0 := datamath.FibonacciNthTerm(term0)
	result1 := datamath.FibonacciNthTerm(term1)
	result2 := datamath.FibonacciNthTerm(term2)
	result3 := datamath.FibonacciNthTerm(term3)
	result8 := datamath.FibonacciNthTerm(term8)
	result25 := datamath.FibonacciNthTerm(term25)

	resultN0 := datamath.FibonacciNthTerm(termN0)
	resultN1 := datamath.FibonacciNthTerm(termN1)
	resultN2 := datamath.FibonacciNthTerm(termN2)
	resultN3 := datamath.FibonacciNthTerm(termN3)
	resultN8 := datamath.FibonacciNthTerm(termN8)
	resultN25 := datamath.FibonacciNthTerm(termN25)

	if result0 != expected0 {
		t.Errorf(
			"got: %v\nwant: %v",
			result0, expected0,
		)
	}
	if result1 != expected1 {
		t.Errorf(
			"got: %v\nwant: %v",
			result1, expected1,
		)
	}
	if result2 != expected2 {
		t.Errorf(
			"got: %v\nwant: %v",
			result2, expected2,
		)
	}
	if result3 != expected3 {
		t.Errorf(
			"got: %v\nwant: %v",
			result3, expected3,
		)
	}
	if result8 != expected8 {
		t.Errorf(
			"got: %v\nwant: %v",
			result8, expected8,
		)
	}
	if result25 != expected25 {
		t.Errorf(
			"got: %v\nwant: %v",
			result25, expected25,
		)
	}

	if resultN0 != expectedN0 {
		t.Errorf(
			"got: %v\nwant: %v",
			resultN0, expectedN0,
		)
	}
	if resultN1 != expectedN1 {
		t.Errorf(
			"got: %v\nwant: %v",
			resultN1, expectedN1,
		)
	}
	if resultN2 != expectedN2 {
		t.Errorf(
			"got: %v\nwant: %v",
			resultN2, expectedN2,
		)
	}
	if resultN3 != expectedN3 {
		t.Errorf(
			"got: %v\nwant: %v",
			resultN3, expectedN3,
		)
	}
	if resultN8 != expectedN8 {
		t.Errorf(
			"got: %v\nwant: %v",
			resultN8, expectedN8,
		)
	}
	if resultN25 != expectedN25 {
		t.Errorf(
			"got: %v\nwant: %v",
			resultN25, expectedN25,
		)
	}
}
