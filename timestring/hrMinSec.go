// Package timestring returns time in human-readable format
package timestring

import "strconv"

// HourMinuteSecond takes seconds as input and returns the calculated
// time in string: x hour(s) y minute(s) z second(s)
func HourMinuteSecond(timeInSecond uint64) string {
	if timeInSecond == 0 {
		return second(timeInSecond)
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
		timeInSecondStr = second(timeInSecond)
		return timeInSecondStr
	}

	// minutes
	timeInMinute = timeInSecond / 60

	remainingSecond = timeInSecond % 60
	if remainingSecond > 0 {
		remainingSecondStr = second(remainingSecond)
	}

	if timeInMinute < 60 {
		timeInMinuteStr = minute(timeInMinute)

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
		remainingMinuteStr = minute(remainingMinute)
	}

	timeInHourStr = hour(timeInHour)

	totalTimeStr := timeInHourStr
	if remainingMinute > 0 {
		totalTimeStr += " " + remainingMinuteStr
	}
	if remainingSecond > 0 {
		totalTimeStr += " " + remainingSecondStr
	}

	return totalTimeStr
}

func second(timeInSecond uint64) string {
	result := strconv.FormatUint(timeInSecond, 10)

	if timeInSecond <= 1 {
		result += " second"
	}

	if timeInSecond > 1 {
		result += " seconds"
	}

	return result
}

func minute(timeInMinute uint64) string {
	result := strconv.FormatUint(timeInMinute, 10)

	if timeInMinute == 1 {
		result += " minute"
	}

	if timeInMinute > 1 {
		result += " minutes"
	}

	return result
}

func hour(timeInHour uint64) string {
	result := strconv.FormatUint(timeInHour, 10)

	if timeInHour == 1 {
		result += " hour"
	}

	if timeInHour > 1 {
		result += " hours"
	}

	return result
}
