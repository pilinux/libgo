// Package string implements functions to manipulate UTF-8 encoded strings
package string

import (
	"bytes"
)

// InsertNewCharNthPosition inserts a new UTF-8 character
// after every N-th character and receives the result as
// a UTF-8 formatted string
//
// Example:
// input := "12345678123456781234567812345678"
// result := InsertNewCharNthPosition(input, 8, '-')
// result will be "12345678-12345678-12345678-12345678"
func InsertNewCharNthPosition(s string, n int, r rune) string {
	var buffer bytes.Buffer
	n1 := n - 1
	l1 := len(s) - 1

	for i, rune := range s {
		buffer.WriteRune(rune)

		if i%n == n1 && i != l1 {
			buffer.WriteRune(r)
		}
	}

	return buffer.String()
}
