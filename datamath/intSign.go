// Package datamath provides mathematical functions which are not implemented in the standard library.
package datamath

// IntSign returns false if the integer is 0 or positive,
// and returns true if the integer is negative.
func IntSign(x int) bool {
	return x < 0
}
