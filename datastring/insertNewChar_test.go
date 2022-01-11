package datastring_test

import (
	"testing"

	"github.com/pilinux/libgo/datastring"
)

func TestInsertNewCharNthPosition(t *testing.T) {
	input := "12345678123456781234567812345678"
	result := datastring.InsertNewCharNthPosition(input, 8, '-')
	expected := "12345678-12345678-12345678-12345678"

	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}
}
