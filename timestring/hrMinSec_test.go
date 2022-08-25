package timestring_test

import (
	"testing"

	"github.com/pilinux/libgo/timestring"
)

func TestHourMinuteSecond(t *testing.T) {
	var result string
	var expected string

	// 0 second
	result = timestring.HourMinuteSecond(0)
	expected = "0 second"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 1 second
	result = timestring.HourMinuteSecond(1)
	expected = "1 second"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 59 seconds
	result = timestring.HourMinuteSecond(59)
	expected = "59 seconds"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 1 minute
	result = timestring.HourMinuteSecond(60)
	expected = "1 minute"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 1 minute 1 second
	result = timestring.HourMinuteSecond(61)
	expected = "1 minute 1 second"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 1 minute 59 seconds
	result = timestring.HourMinuteSecond(119)
	expected = "1 minute 59 seconds"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 2 minutes
	result = timestring.HourMinuteSecond(120)
	expected = "2 minutes"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 2 minutes 1 second
	result = timestring.HourMinuteSecond(121)
	expected = "2 minutes 1 second"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 2 minutes 59 seconds
	result = timestring.HourMinuteSecond(179)
	expected = "2 minutes 59 seconds"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 59 minutes
	result = timestring.HourMinuteSecond(3540)
	expected = "59 minutes"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 59 minutes 1 second
	result = timestring.HourMinuteSecond(3541)
	expected = "59 minutes 1 second"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 59 minutes 59 seconds
	result = timestring.HourMinuteSecond(3599)
	expected = "59 minutes 59 seconds"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 1 hour
	result = timestring.HourMinuteSecond(3600)
	expected = "1 hour"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 1 hour 1 second
	result = timestring.HourMinuteSecond(3601)
	expected = "1 hour 1 second"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 1 hour 59 seconds
	result = timestring.HourMinuteSecond(3659)
	expected = "1 hour 59 seconds"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 1 hour 1 minute
	result = timestring.HourMinuteSecond(3660)
	expected = "1 hour 1 minute"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 1 hour 1 minute 1 second
	result = timestring.HourMinuteSecond(3661)
	expected = "1 hour 1 minute 1 second"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 1 hour 1 minute 59 seconds
	result = timestring.HourMinuteSecond(3719)
	expected = "1 hour 1 minute 59 seconds"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 1 hour 2 minutes
	result = timestring.HourMinuteSecond(3720)
	expected = "1 hour 2 minutes"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 1 hour 2 minutes 1 second
	result = timestring.HourMinuteSecond(3721)
	expected = "1 hour 2 minutes 1 second"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 1 hour 59 minutes 59 seconds
	result = timestring.HourMinuteSecond(7199)
	expected = "1 hour 59 minutes 59 seconds"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 2 hours
	result = timestring.HourMinuteSecond(7200)
	expected = "2 hours"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 2 hours 1 second
	result = timestring.HourMinuteSecond(7201)
	expected = "2 hours 1 second"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 2 hours 59 seconds
	result = timestring.HourMinuteSecond(7259)
	expected = "2 hours 59 seconds"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 2 hours 1 minute
	result = timestring.HourMinuteSecond(7260)
	expected = "2 hours 1 minute"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 2 hours 1 minute 1 second
	result = timestring.HourMinuteSecond(7261)
	expected = "2 hours 1 minute 1 second"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 2 hours 1 minute 59 seconds
	result = timestring.HourMinuteSecond(7319)
	expected = "2 hours 1 minute 59 seconds"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 2 hours 2 minutes
	result = timestring.HourMinuteSecond(7320)
	expected = "2 hours 2 minutes"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 2 hours 2 minutes 1 second
	result = timestring.HourMinuteSecond(7321)
	expected = "2 hours 2 minutes 1 second"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 2 hours 59 minutes 59 seconds
	result = timestring.HourMinuteSecond(10799)
	expected = "2 hours 59 minutes 59 seconds"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 24 hours
	result = timestring.HourMinuteSecond(86400)
	expected = "24 hours"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 48 hours
	result = timestring.HourMinuteSecond(172800)
	expected = "48 hours"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}

	// 72 hours
	result = timestring.HourMinuteSecond(259200)
	expected = "72 hours"
	if result != expected {
		t.Errorf(
			"got: %v\nwant: %v",
			result, expected,
		)
	}
}
