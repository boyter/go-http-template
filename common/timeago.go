// taken from https://github.com/justincampbell/timeago
// SPDX-License-Identifier: MIT

package common

import (
	"bytes"
	"fmt"
	"math"
	"time"
)

const (
	minute = 1
	hour   = minute * 60
	day    = hour * 24
	month  = day * 30
	year   = day * 365
)

// FromDuration returns a friendly string representing an approximation of the
// given duration
func FromDuration(d time.Duration) string {
	seconds := round(d.Seconds())

	if seconds < 30 {
		return "less than a minute"
	}

	if seconds < 90 {
		return "1 minute"
	}

	minutes := div(seconds, 60)

	if minutes < 45 {
		return fmt.Sprintf("%0d minutes", minutes)
	}

	hours := div(minutes, 60)

	if minutes < day {
		return fmt.Sprintf("about %s", pluralize(hours, "hour"))
	}

	if minutes < (42 * hour) {
		return "1 day"
	}

	days := div(hours, 24)

	if minutes < (30 * day) {
		return pluralize(days, "day")
	}

	months := div(days, 30)

	if minutes < (45 * day) {
		return "about 1 month"
	}

	if minutes < (60 * day) {
		return "about 2 months"
	}

	if minutes < year {
		return pluralize(months, "month")
	}

	rem := minutes % year
	years := minutes / year

	if rem < (3 * month) {
		return fmt.Sprintf("about %s", pluralize(years, "year"))
	}
	if rem < (9 * month) {
		return fmt.Sprintf("over %s", pluralize(years, "year"))
	}

	years++
	return fmt.Sprintf("almost %s", pluralize(years, "year"))
}

// FromTime returns a friendly string representing the approximate difference
// from the given time and time.Now()
func FromTime(t time.Time) string {
	now := time.Now()

	var d time.Duration
	var suffix string

	if t.Before(now) {
		d = now.Sub(t)
		suffix = "ago"
	} else {
		d = t.Sub(now)
		suffix = "from now"
	}

	return fmt.Sprintf("%s %s", FromDuration(d), suffix)
}

func pluralize(i int, s string) string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d %s", i, s))
	if i != 1 {
		buf.WriteString("s")
	}
	return buf.String()
}

func round(f float64) int {
	return int(math.Floor(f + .50))
}

func div(numerator int, denominator int) int {
	rem := numerator % denominator
	result := numerator / denominator

	if rem >= (denominator / 2) {
		result++
	}

	return result
}
