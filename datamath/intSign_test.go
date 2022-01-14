// Package datamath provides mathematical functions which are not implemented in the standard library.
package datamath_test

import (
	"testing"

	"github.com/pilinux/libgo/datamath"
)

func TestIntSign(t *testing.T) {
	input1 := 0
	input2 := -0
	input3 := 1
	input4 := -1
	input5 := 9223372036854775807  // max value of 64-bit int
	input6 := -9223372036854775808 // min value of 64-bit int

	expected1 := false
	expected2 := false
	expected3 := false
	expected4 := true
	expected5 := false
	expected6 := true

	result1 := datamath.IntSign(input1)
	result2 := datamath.IntSign(input2)
	result3 := datamath.IntSign(input3)
	result4 := datamath.IntSign(input4)
	result5 := datamath.IntSign(input5)
	result6 := datamath.IntSign(input6)

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

	if result4 != expected4 {
		t.Errorf(
			"got: %v\nwant: %v",
			result4, expected4,
		)
	}

	if result5 != expected5 {
		t.Errorf(
			"got: %v\nwant: %v",
			result5, expected5,
		)
	}

	if result6 != expected6 {
		t.Errorf(
			"got: %v\nwant: %v",
			result6, expected6,
		)
	}
}
