package timestring_test

import (
	"fmt"
	"testing"

	"github.com/pilinux/libgo/timestring"
)

type hourMinuteSecondTest struct {
	input    uint64
	expected string
}

var hourMinuteSecondTests = []hourMinuteSecondTest{
	{0, "0 second"},
	{1, "1 second"},
	{59, "59 seconds"},
	{60, "1 minute"},
	{61, "1 minute 1 second"},
	{119, "1 minute 59 seconds"},
	{120, "2 minutes"},
	{121, "2 minutes 1 second"},
	{179, "2 minutes 59 seconds"},
	{3540, "59 minutes"},
	{3541, "59 minutes 1 second"},
	{3599, "59 minutes 59 seconds"},
	{3600, "1 hour"},
	{3601, "1 hour 1 second"},
	{3659, "1 hour 59 seconds"},
	{3660, "1 hour 1 minute"},
	{3661, "1 hour 1 minute 1 second"},
	{3719, "1 hour 1 minute 59 seconds"},
	{3720, "1 hour 2 minutes"},
	{3721, "1 hour 2 minutes 1 second"},
	{7199, "1 hour 59 minutes 59 seconds"},
	{7200, "2 hours"},
	{7201, "2 hours 1 second"},
	{7259, "2 hours 59 seconds"},
	{7260, "2 hours 1 minute"},
	{7261, "2 hours 1 minute 1 second"},
	{7319, "2 hours 1 minute 59 seconds"},
	{7320, "2 hours 2 minutes"},
	{7321, "2 hours 2 minutes 1 second"},
	{10799, "2 hours 59 minutes 59 seconds"},
	{86400, "24 hours"},
	{172800, "48 hours"},
	{259200, "72 hours"},
}

func TestHourMinuteSecond(t *testing.T) {
	for _, test := range hourMinuteSecondTests {
		result := timestring.HourMinuteSecond(test.input)
		expected := test.expected
		if result != expected {
			t.Errorf(
				"got: %v\nwant: %v",
				result, expected,
			)
		}
	}
}

func ExampleHourMinuteSecond() {
	fmt.Println(timestring.HourMinuteSecond(10799))
	// Output: 2 hours 59 minutes 59 seconds
}
