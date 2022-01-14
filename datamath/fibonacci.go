// Package datamath provides mathematical functions which are not implemented in the standard library.
package datamath

// FibonacciNthTerm returns the n-th term from the Fibonacci sequence.
// The Fibonacci sequence is 0 (0th term), 1 (1st term), 1 (2nd term), 2 (3rd term), 3 (4th term),
// 5 (5th term), 8 (6th term), 13 (7th term), 21 (8th term), 34 (9th term), ...
// Example: FibonacciNthTerm(8) = 21
//
// FibonacciNthTerm extends the Fibonacci numbers to negative integers also.
// Example: FibonacciNthTerm(-8) = -21
func FibonacciNthTerm(x int) int {
	signX := IntSign(x)

	if signX {
		x = x * (-1)
	}

	n := make([]int, x+1, x+2)

	if x < 2 {
		n = n[0:2]
	}
	n[0] = 0
	n[1] = 1

	for i := 2; i <= x; i++ {
		n[i] = n[i-1] + n[i-2]
	}

	if signX {
		n[x] = n[x] * (-1)
	}

	return n[x]
}

// Super inefficient approach
/*
func FibonacciNthTerm(x int) int {
	signX := IntSign(x)

	if signX {
		x = x * (-1)
	}

	if x <= 1 {
		if signX {
			x = x * (-1)
		}
		return x
	}

	nthTerm := FibonacciNthTerm(x-1) + FibonacciNthTerm(x-2)
	if signX {
		nthTerm = nthTerm * (-1)
	}
	return nthTerm
}
*/
