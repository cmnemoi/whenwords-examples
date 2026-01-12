# whenwords for Go

Human-friendly time formatting and parsing.

## Installation

Copy the `whenwords.go` file into your project, or import it as a local module:

```go
import "path/to/whenwords"
```

For local development, the module name is `whenwords` as specified in `go.mod`.

## Quick start

```go
package main

import (
	"fmt"
	"time"
	"whenwords"
)

func main() {
	now := time.Now().Unix()
	pastTime := now - 3600 // 1 hour ago

	// Relative time
	fmt.Println(whenwords.Timeago(pastTime, now))  // "1 hour ago"

	// Duration formatting
	duration, _ := whenwords.Duration(3661, nil)
	fmt.Println(duration)  // "1 hour, 1 minute"

	// Parse human-written durations
	seconds, _ := whenwords.ParseDuration("2h 30m")
	fmt.Println(seconds)  // 9000

	// Contextual dates
	fmt.Println(whenwords.HumanDate(pastTime, now))  // "Today"

	// Date ranges
	start := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC).Unix()
	end := time.Date(2024, 1, 22, 0, 0, 0, 0, time.UTC).Unix()
	fmt.Println(whenwords.DateRange(start, end))  // "January 15–22, 2024"
}
```

## Functions

### Timeago(timestamp, reference int64) → string

Returns a human-readable relative time string.

**Parameters:**

- `timestamp`: Unix timestamp in seconds (the time to describe)
- `reference`: Unix timestamp in seconds (the "current" time to compare against)

**Returns:** A string like "3 hours ago" or "in 2 days"

**Examples:**

```go
now := time.Now().Unix()
past := now - 7200

whenwords.Timeago(past, now)  // "2 hours ago"
whenwords.Timeago(now + 3600, now)  // "in 1 hour"
whenwords.Timeago(now, now)  // "just now"
```

### Duration(seconds float64, opts \*DurationOptions) → (string, error)

Formats a duration in seconds as a human-readable string.

**Parameters:**

- `seconds`: Non-negative number of seconds
- `opts`: Optional configuration:
  - `Compact`: If true, use "2h 30m" style (default: false)
  - `MaxUnits`: Maximum number of units to show (default: 2)

**Returns:** Formatted duration string and error (if any)

**Examples:**

```go
whenwords.Duration(3661, nil)  // "1 hour, 1 minute", nil
whenwords.Duration(3661, &whenwords.DurationOptions{Compact: true})  // "1h 1m", nil
whenwords.Duration(93661, &whenwords.DurationOptions{MaxUnits: 3})  // "1 day, 2 hours, 1 minute", nil
whenwords.Duration(-100, nil)  // "", error
```

### ParseDuration(s string) → (int64, error)

Parses a human-written duration string into seconds.

**Parameters:**

- `s`: Duration string in various formats

**Accepted formats:**

- Compact: "2h30m", "2h 30m", "2h, 30m"
- Verbose: "2 hours 30 minutes", "2 hours and 30 minutes"
- Decimal: "2.5 hours", "1.5h"
- Colon notation: "2:30" (h:mm), "1:30:00" (h:mm:ss)

**Unit aliases:**

- seconds: s, sec, secs, second, seconds
- minutes: m, min, mins, minute, minutes
- hours: h, hr, hrs, hour, hours
- days: d, day, days
- weeks: w, wk, wks, week, weeks

**Returns:** Number of seconds and error (if any)

**Examples:**

```go
whenwords.ParseDuration("2h 30m")  // 9000, nil
whenwords.ParseDuration("2 hours and 30 minutes")  // 9000, nil
whenwords.ParseDuration("2:30")  // 9000, nil
whenwords.ParseDuration("1.5h")  // 5400, nil
whenwords.ParseDuration("invalid")  // 0, error
```

### HumanDate(timestamp, reference int64) → string

Returns a contextual date string relative to a reference date.

**Parameters:**

- `timestamp`: Unix timestamp in seconds (the date to format)
- `reference`: Unix timestamp in seconds (the "current" date for comparison)

**Returns:** Contextual string like "Today", "Yesterday", "Last Friday", or "March 15"

**Behavior:**

- Same day → "Today"
- Previous day → "Yesterday"
- Next day → "Tomorrow"
- Within past 2-6 days → "Last {weekday}"
- Within next 2-6 days → "This {weekday}"
- Same year → "{Month} {day}"
- Different year → "{Month} {day}, {year}"

**Examples:**

```go
now := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC).Unix()
today := now
yesterday := now - 86400
lastWeek := now - 3*86400

whenwords.HumanDate(today, now)  // "Today"
whenwords.HumanDate(yesterday, now)  // "Yesterday"
whenwords.HumanDate(lastWeek, now)  // "Last Friday"
```

### DateRange(start, end int64) → string

Formats a date range with smart abbreviation.

**Parameters:**

- `start`: Unix timestamp in seconds (start of range)
- `end`: Unix timestamp in seconds (end of range)

**Returns:** Formatted date range string

**Behavior:**

- Same day: "January 15, 2024"
- Same month: "January 15–22, 2024"
- Same year: "January 15 – February 7, 2024"
- Different years: "December 28, 2023 – January 3, 2024"

**Edge cases:**

- If `start` > `end`, they are automatically swapped

**Examples:**

```go
jan15 := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC).Unix()
jan22 := time.Date(2024, 1, 22, 0, 0, 0, 0, time.UTC).Unix()
feb15 := time.Date(2024, 2, 15, 0, 0, 0, 0, time.UTC).Unix()

whenwords.DateRange(jan15, jan15)  // "January 15, 2024"
whenwords.DateRange(jan15, jan22)  // "January 15–22, 2024"
whenwords.DateRange(jan15, feb15)  // "January 15 – February 15, 2024"
```

## Error handling

Functions that can error return Go's standard `(value, error)` tuple:

```go
result, err := whenwords.Duration(-100, nil)
if err != nil {
	// Handle error: "invalid seconds: must be non-negative and finite"
}

seconds, err := whenwords.ParseDuration("")
if err != nil {
	// Handle error: "empty string"
}
```

**Error conditions:**

- `Duration`: Negative seconds, NaN, or infinite values
- `ParseDuration`: Empty string, no parseable units, negative values

Functions that cannot error (`Timeago`, `HumanDate`, `DateRange`) return strings directly.

## Accepted types

All functions accept Unix timestamps as `int64` values (seconds since 1970-01-01 UTC).

To convert Go's `time.Time` to Unix seconds:

```go
t := time.Now()
timestamp := t.Unix()
```

To convert Unix seconds to `time.Time`:

```go
timestamp := int64(1704067200)
t := time.Unix(timestamp, 0).UTC()
```

For `Duration`, the input is a `float64` representing seconds, allowing for fractional seconds if needed.

For `ParseDuration`, the input is a `string` that gets parsed into `int64` seconds.
