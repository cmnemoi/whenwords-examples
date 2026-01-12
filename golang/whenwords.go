package whenwords

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// DurationOptions holds options for duration formatting
type DurationOptions struct {
	Compact  bool
	MaxUnits int
}

// Timeago returns a human-readable relative time string
func Timeago(timestamp, reference int64) string {
	diff := reference - timestamp
	absDiff := diff
	if absDiff < 0 {
		absDiff = -absDiff
	}

	var n int64
	var unit string

	// Determine the appropriate unit and value
	if absDiff < 45 {
		return "just now"
	} else if absDiff < 90 {
		n = 1
		unit = "minute"
	} else if absDiff < 45*60 {
		n = int64(math.Round(float64(absDiff) / 60.0))
		unit = "minute"
	} else if absDiff < 90*60 {
		n = 1
		unit = "hour"
	} else if absDiff < 22*3600 {
		n = int64(math.Round(float64(absDiff) / 3600.0))
		unit = "hour"
	} else if absDiff < 36*3600 {
		n = 1
		unit = "day"
	} else if absDiff < 26*86400 {
		n = int64(math.Round(float64(absDiff) / 86400.0))
		unit = "day"
	} else if absDiff < 46*86400 {
		n = 1
		unit = "month"
	} else if absDiff < 320*86400 {
		n = int64(math.Round(float64(absDiff) / (30.0 * 86400.0)))
		// Cap at 10 months to avoid "11 months ago" right before "1 year ago"
		if n > 10 {
			n = 10
		}
		unit = "month"
	} else if absDiff < 548*86400 {
		n = 1
		unit = "year"
	} else {
		n = int64(math.Round(float64(absDiff) / (365.0 * 86400.0)))
		unit = "year"
	}

	// Pluralize if needed
	if n != 1 {
		unit += "s"
	}

	// Format as past or future
	if diff > 0 {
		return fmt.Sprintf("%d %s ago", n, unit)
	} else if diff < 0 {
		return fmt.Sprintf("in %d %s", n, unit)
	}
	return "just now"
}

// Duration formats a duration in seconds as a human-readable string
func Duration(seconds float64, opts *DurationOptions) (string, error) {
	if seconds < 0 || math.IsNaN(seconds) || math.IsInf(seconds, 0) {
		return "", errors.New("invalid seconds: must be non-negative and finite")
	}

	// Set default options
	compact := false
	maxUnits := 2
	if opts != nil {
		compact = opts.Compact
		if opts.MaxUnits > 0 {
			maxUnits = opts.MaxUnits
		}
	}

	if seconds == 0 {
		if compact {
			return "0s", nil
		}
		return "0 seconds", nil
	}

	// Unit definitions (in seconds)
	type unit struct {
		name     string
		shortcut string
		seconds  float64
	}
	units := []unit{
		{"year", "y", 365 * 24 * 3600},
		{"month", "mo", 30 * 24 * 3600},
		{"day", "d", 24 * 3600},
		{"hour", "h", 3600},
		{"minute", "m", 60},
		{"second", "s", 1},
	}

	// Calculate each unit
	remaining := seconds
	var parts []string
	for _, u := range units {
		if remaining >= u.seconds {
			count := math.Floor(remaining / u.seconds)
			remaining = math.Mod(remaining, u.seconds)

			if compact {
				parts = append(parts, fmt.Sprintf("%.0f%s", count, u.shortcut))
			} else {
				unitName := u.name
				if count != 1 {
					unitName += "s"
				}
				parts = append(parts, fmt.Sprintf("%.0f %s", count, unitName))
			}

			if len(parts) >= maxUnits {
				break
			}
		}
	}

	// Join parts
	if compact {
		return strings.Join(parts, " "), nil
	}
	return strings.Join(parts, ", "), nil
}

// ParseDuration parses a human-written duration string into seconds
func ParseDuration(s string) (int64, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, errors.New("empty string")
	}

	// Check for colon notation (h:mm or h:mm:ss)
	colonPattern := regexp.MustCompile(`^(\d+):(\d+)(?::(\d+))?$`)
	if match := colonPattern.FindStringSubmatch(s); match != nil {
		hours, _ := strconv.ParseInt(match[1], 10, 64)
		minutes, _ := strconv.ParseInt(match[2], 10, 64)
		seconds := int64(0)
		if match[3] != "" {
			seconds, _ = strconv.ParseInt(match[3], 10, 64)
		}
		total := hours*3600 + minutes*60 + seconds
		if total < 0 {
			return 0, errors.New("negative duration")
		}
		return total, nil
	}

	// Normalize string: lowercase, remove extra whitespace
	s = strings.ToLower(s)
	s = regexp.MustCompile(`\s+`).ReplaceAllString(s, " ")

	// Remove commas and "and"
	s = strings.ReplaceAll(s, ",", " ")
	s = strings.ReplaceAll(s, " and ", " ")
	s = strings.TrimSpace(s)

	// Pattern to match number + unit
	pattern := regexp.MustCompile(`([+-]?\d+(?:\.\d+)?)\s*([a-z]+)`)
	matches := pattern.FindAllStringSubmatch(s, -1)

	if len(matches) == 0 {
		return 0, errors.New("no parseable units found")
	}

	var total float64
	for _, match := range matches {
		valueStr := match[1]
		unit := match[2]

		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid number: %s", valueStr)
		}

		if value < 0 {
			return 0, errors.New("negative duration")
		}

		// Map unit to seconds
		var multiplier float64
		switch unit {
		case "s", "sec", "secs", "second", "seconds":
			multiplier = 1
		case "m", "min", "mins", "minute", "minutes":
			multiplier = 60
		case "h", "hr", "hrs", "hour", "hours":
			multiplier = 3600
		case "d", "day", "days":
			multiplier = 86400
		case "w", "wk", "wks", "week", "weeks":
			multiplier = 604800
		default:
			return 0, fmt.Errorf("unknown unit: %s", unit)
		}

		total += value * multiplier
	}

	if total < 0 {
		return 0, errors.New("negative duration")
	}

	return int64(total), nil
}

// HumanDate returns a contextual date string
func HumanDate(timestamp, reference int64) string {
	// Convert to time.Time in UTC
	ts := time.Unix(timestamp, 0).UTC()
	ref := time.Unix(reference, 0).UTC()

	// Get date components (ignore time)
	tsYear, tsMonth, tsDay := ts.Date()
	refYear, refMonth, refDay := ref.Date()

	// Check if same day
	if tsYear == refYear && tsMonth == refMonth && tsDay == refDay {
		return "Today"
	}

	// Check if yesterday
	yesterday := ref.AddDate(0, 0, -1)
	yYear, yMonth, yDay := yesterday.Date()
	if tsYear == yYear && tsMonth == yMonth && tsDay == yDay {
		return "Yesterday"
	}

	// Check if tomorrow
	tomorrow := ref.AddDate(0, 0, 1)
	tYear, tMonth, tDay := tomorrow.Date()
	if tsYear == tYear && tsMonth == tMonth && tsDay == tDay {
		return "Tomorrow"
	}

	// Calculate day difference
	daysDiff := int(ts.Sub(ref).Hours() / 24)
	if ts.Before(ref) {
		// Past dates
		if daysDiff >= -6 && daysDiff <= -2 {
			return "Last " + ts.Weekday().String()
		}
	} else {
		// Future dates
		if daysDiff >= 2 && daysDiff <= 6 {
			return "This " + ts.Weekday().String()
		}
	}

	// Same year: "Month Day"
	if tsYear == refYear {
		return ts.Format("January 2")
	}

	// Different year: "Month Day, Year"
	return ts.Format("January 2, 2006")
}

// DateRange formats a date range with smart abbreviation
func DateRange(start, end int64) string {
	// Swap if necessary
	if start > end {
		start, end = end, start
	}

	startTime := time.Unix(start, 0).UTC()
	endTime := time.Unix(end, 0).UTC()

	startYear, startMonth, startDay := startTime.Date()
	endYear, endMonth, endDay := endTime.Date()

	// Same day
	if startYear == endYear && startMonth == endMonth && startDay == endDay {
		return startTime.Format("January 2, 2006")
	}

	// Same month and year
	if startYear == endYear && startMonth == endMonth {
		return fmt.Sprintf("%s %d–%d, %d",
			startMonth.String(), startDay, endDay, startYear)
	}

	// Same year, different months
	if startYear == endYear {
		return fmt.Sprintf("%s %d – %s %d, %d",
			startMonth.String(), startDay,
			endMonth.String(), endDay,
			startYear)
	}

	// Different years
	return fmt.Sprintf("%s %d, %d – %s %d, %d",
		startMonth.String(), startDay, startYear,
		endMonth.String(), endDay, endYear)
}
