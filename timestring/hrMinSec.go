// Package timestring returns time in human-readable format
package timestring

import "strconv"

// HourMinuteSecond takes seconds as input and returns the calculated
// time in string: x hour(s) y minute(s) z second(s)
func HourMinuteSecond(timeInSecond uint64) string {
	if timeInSecond == 0 {
		return strconv.FormatUint(timeInSecond, 10) + " second"
	}

	var timeInMinute uint64
	var timeInHour uint64
	var remainingSecond uint64
	var remainingMinute uint64

	var timeInSecondStr string
	var timeInMinuteStr string
	var timeInHourStr string
	var remainingSecondStr string
	var remainingMinuteStr string

	// seconds
	if timeInSecond < 60 {
		timeInSecondStr = strconv.FormatUint(timeInSecond, 10)
		if timeInSecond == 1 {
			timeInSecondStr += " second"
		}
		if timeInSecond > 1 {
			timeInSecondStr += " seconds"
		}
		return timeInSecondStr
	}

	// minutes
	timeInMinute = timeInSecond / 60

	remainingSecond = timeInSecond % 60
	if remainingSecond > 0 {
		remainingSecondStr = strconv.FormatUint(remainingSecond, 10)
		if remainingSecond == 1 {
			remainingSecondStr += " second"
		}
		if remainingSecond > 1 {
			remainingSecondStr += " seconds"
		}
	}

	if timeInMinute < 60 {
		timeInMinuteStr = strconv.FormatUint(timeInMinute, 10)
		if timeInMinute == 1 {
			timeInMinuteStr += " minute"
		}
		if timeInMinute > 1 {
			timeInMinuteStr += " minutes"
		}

		totalTimeStr := timeInMinuteStr
		if remainingSecond > 0 {
			totalTimeStr += " " + remainingSecondStr
		}

		return totalTimeStr
	}

	// hours
	timeInHour = timeInMinute / 60
	remainingMinute = timeInMinute % 60
	if remainingMinute > 0 {
		remainingMinuteStr = strconv.FormatUint(remainingMinute, 10)
		if remainingMinute == 1 {
			remainingMinuteStr += " minute"
		}
		if remainingMinute > 1 {
			remainingMinuteStr += " minutes"
		}
	}

	timeInHourStr = strconv.FormatUint(timeInHour, 10)
	if timeInHour == 1 {
		timeInHourStr += " hour"
	}
	if timeInHour > 1 {
		timeInHourStr += " hours"
	}

	totalTimeStr := timeInHourStr
	if remainingMinute > 0 {
		totalTimeStr += " " + remainingMinuteStr
	}
	if remainingSecond > 0 {
		totalTimeStr += " " + remainingSecondStr
	}

	return totalTimeStr
}
