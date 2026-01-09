package whenwords

import (
	"testing"
)

// timeago tests
func TestTimeagoJustNowIdenticalTimestamps(t *testing.T) {
	result := Timeago(1704067200, 1704067200)
	if result != "just now" {
		t.Errorf("Expected 'just now', got '%s'", result)
	}
}

func TestTimeagoJustNow30SecondsAgo(t *testing.T) {
	result := Timeago(1704067170, 1704067200)
	if result != "just now" {
		t.Errorf("Expected 'just now', got '%s'", result)
	}
}

func TestTimeagoJustNow44SecondsAgo(t *testing.T) {
	result := Timeago(1704067156, 1704067200)
	if result != "just now" {
		t.Errorf("Expected 'just now', got '%s'", result)
	}
}

func TestTimeago1MinuteAgo45Seconds(t *testing.T) {
	result := Timeago(1704067155, 1704067200)
	if result != "1 minute ago" {
		t.Errorf("Expected '1 minute ago', got '%s'", result)
	}
}

func TestTimeago1MinuteAgo89Seconds(t *testing.T) {
	result := Timeago(1704067111, 1704067200)
	if result != "1 minute ago" {
		t.Errorf("Expected '1 minute ago', got '%s'", result)
	}
}

func TestTimeago2MinutesAgo90Seconds(t *testing.T) {
	result := Timeago(1704067110, 1704067200)
	if result != "2 minutes ago" {
		t.Errorf("Expected '2 minutes ago', got '%s'", result)
	}
}

func TestTimeago30MinutesAgo(t *testing.T) {
	result := Timeago(1704065400, 1704067200)
	if result != "30 minutes ago" {
		t.Errorf("Expected '30 minutes ago', got '%s'", result)
	}
}

func TestTimeago44MinutesAgo(t *testing.T) {
	result := Timeago(1704064560, 1704067200)
	if result != "44 minutes ago" {
		t.Errorf("Expected '44 minutes ago', got '%s'", result)
	}
}

func TestTimeago1HourAgo45Minutes(t *testing.T) {
	result := Timeago(1704064500, 1704067200)
	if result != "1 hour ago" {
		t.Errorf("Expected '1 hour ago', got '%s'", result)
	}
}

func TestTimeago1HourAgo89Minutes(t *testing.T) {
	result := Timeago(1704061860, 1704067200)
	if result != "1 hour ago" {
		t.Errorf("Expected '1 hour ago', got '%s'", result)
	}
}

func TestTimeago2HoursAgo90Minutes(t *testing.T) {
	result := Timeago(1704061800, 1704067200)
	if result != "2 hours ago" {
		t.Errorf("Expected '2 hours ago', got '%s'", result)
	}
}

func TestTimeago5HoursAgo(t *testing.T) {
	result := Timeago(1704049200, 1704067200)
	if result != "5 hours ago" {
		t.Errorf("Expected '5 hours ago', got '%s'", result)
	}
}

func TestTimeago21HoursAgo(t *testing.T) {
	result := Timeago(1703991600, 1704067200)
	if result != "21 hours ago" {
		t.Errorf("Expected '21 hours ago', got '%s'", result)
	}
}

func TestTimeago1DayAgo22Hours(t *testing.T) {
	result := Timeago(1703988000, 1704067200)
	if result != "1 day ago" {
		t.Errorf("Expected '1 day ago', got '%s'", result)
	}
}

func TestTimeago1DayAgo35Hours(t *testing.T) {
	result := Timeago(1703941200, 1704067200)
	if result != "1 day ago" {
		t.Errorf("Expected '1 day ago', got '%s'", result)
	}
}

func TestTimeago2DaysAgo36Hours(t *testing.T) {
	result := Timeago(1703937600, 1704067200)
	if result != "2 days ago" {
		t.Errorf("Expected '2 days ago', got '%s'", result)
	}
}

func TestTimeago7DaysAgo(t *testing.T) {
	result := Timeago(1703462400, 1704067200)
	if result != "7 days ago" {
		t.Errorf("Expected '7 days ago', got '%s'", result)
	}
}

func TestTimeago25DaysAgo(t *testing.T) {
	result := Timeago(1701907200, 1704067200)
	if result != "25 days ago" {
		t.Errorf("Expected '25 days ago', got '%s'", result)
	}
}

func TestTimeago1MonthAgo26Days(t *testing.T) {
	result := Timeago(1701820800, 1704067200)
	if result != "1 month ago" {
		t.Errorf("Expected '1 month ago', got '%s'", result)
	}
}

func TestTimeago1MonthAgo45Days(t *testing.T) {
	result := Timeago(1700179200, 1704067200)
	if result != "1 month ago" {
		t.Errorf("Expected '1 month ago', got '%s'", result)
	}
}

func TestTimeago2MonthsAgo46Days(t *testing.T) {
	result := Timeago(1700092800, 1704067200)
	if result != "2 months ago" {
		t.Errorf("Expected '2 months ago', got '%s'", result)
	}
}

func TestTimeago6MonthsAgo(t *testing.T) {
	result := Timeago(1688169600, 1704067200)
	if result != "6 months ago" {
		t.Errorf("Expected '6 months ago', got '%s'", result)
	}
}

func TestTimeago10MonthsAgo319Days(t *testing.T) {
	result := Timeago(1676505600, 1704067200)
	if result != "10 months ago" {
		t.Errorf("Expected '10 months ago', got '%s'", result)
	}
}

func TestTimeago1YearAgo320Days(t *testing.T) {
	result := Timeago(1676419200, 1704067200)
	if result != "1 year ago" {
		t.Errorf("Expected '1 year ago', got '%s'", result)
	}
}

func TestTimeago1YearAgo547Days(t *testing.T) {
	result := Timeago(1656806400, 1704067200)
	if result != "1 year ago" {
		t.Errorf("Expected '1 year ago', got '%s'", result)
	}
}

func TestTimeago2YearsAgo548Days(t *testing.T) {
	result := Timeago(1656720000, 1704067200)
	if result != "2 years ago" {
		t.Errorf("Expected '2 years ago', got '%s'", result)
	}
}

func TestTimeago5YearsAgo(t *testing.T) {
	result := Timeago(1546300800, 1704067200)
	if result != "5 years ago" {
		t.Errorf("Expected '5 years ago', got '%s'", result)
	}
}

func TestTimeagoFutureInJustNow30Seconds(t *testing.T) {
	result := Timeago(1704067230, 1704067200)
	if result != "just now" {
		t.Errorf("Expected 'just now', got '%s'", result)
	}
}

func TestTimeagoFutureIn1Minute(t *testing.T) {
	result := Timeago(1704067260, 1704067200)
	if result != "in 1 minute" {
		t.Errorf("Expected 'in 1 minute', got '%s'", result)
	}
}

func TestTimeagoFutureIn5Minutes(t *testing.T) {
	result := Timeago(1704067500, 1704067200)
	if result != "in 5 minutes" {
		t.Errorf("Expected 'in 5 minutes', got '%s'", result)
	}
}

func TestTimeagoFutureIn1Hour(t *testing.T) {
	result := Timeago(1704070200, 1704067200)
	if result != "in 1 hour" {
		t.Errorf("Expected 'in 1 hour', got '%s'", result)
	}
}

func TestTimeagoFutureIn3Hours(t *testing.T) {
	result := Timeago(1704078000, 1704067200)
	if result != "in 3 hours" {
		t.Errorf("Expected 'in 3 hours', got '%s'", result)
	}
}

func TestTimeagoFutureIn1Day(t *testing.T) {
	result := Timeago(1704150000, 1704067200)
	if result != "in 1 day" {
		t.Errorf("Expected 'in 1 day', got '%s'", result)
	}
}

func TestTimeagoFutureIn2Days(t *testing.T) {
	result := Timeago(1704240000, 1704067200)
	if result != "in 2 days" {
		t.Errorf("Expected 'in 2 days', got '%s'", result)
	}
}

func TestTimeagoFutureIn1Month(t *testing.T) {
	result := Timeago(1706745600, 1704067200)
	if result != "in 1 month" {
		t.Errorf("Expected 'in 1 month', got '%s'", result)
	}
}

func TestTimeagoFutureIn1Year(t *testing.T) {
	result := Timeago(1735689600, 1704067200)
	if result != "in 1 year" {
		t.Errorf("Expected 'in 1 year', got '%s'", result)
	}
}

// duration tests
func TestDurationZeroSeconds(t *testing.T) {
	result, err := Duration(0, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != "0 seconds" {
		t.Errorf("Expected '0 seconds', got '%s'", result)
	}
}

func TestDuration1Second(t *testing.T) {
	result, err := Duration(1, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != "1 second" {
		t.Errorf("Expected '1 second', got '%s'", result)
	}
}

func TestDuration45Seconds(t *testing.T) {
	result, err := Duration(45, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != "45 seconds" {
		t.Errorf("Expected '45 seconds', got '%s'", result)
	}
}

func TestDuration1Minute(t *testing.T) {
	result, err := Duration(60, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != "1 minute" {
		t.Errorf("Expected '1 minute', got '%s'", result)
	}
}

func TestDuration1Minute30Seconds(t *testing.T) {
	result, err := Duration(90, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != "1 minute, 30 seconds" {
		t.Errorf("Expected '1 minute, 30 seconds', got '%s'", result)
	}
}

func TestDuration2Minutes(t *testing.T) {
	result, err := Duration(120, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != "2 minutes" {
		t.Errorf("Expected '2 minutes', got '%s'", result)
	}
}

func TestDuration1Hour(t *testing.T) {
	result, err := Duration(3600, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != "1 hour" {
		t.Errorf("Expected '1 hour', got '%s'", result)
	}
}

func TestDuration1Hour1Minute(t *testing.T) {
	result, err := Duration(3661, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != "1 hour, 1 minute" {
		t.Errorf("Expected '1 hour, 1 minute', got '%s'", result)
	}
}

func TestDuration1Hour30Minutes(t *testing.T) {
	result, err := Duration(5400, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != "1 hour, 30 minutes" {
		t.Errorf("Expected '1 hour, 30 minutes', got '%s'", result)
	}
}

func TestDuration2Hours30Minutes(t *testing.T) {
	result, err := Duration(9000, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != "2 hours, 30 minutes" {
		t.Errorf("Expected '2 hours, 30 minutes', got '%s'", result)
	}
}

func TestDuration1Day(t *testing.T) {
	result, err := Duration(86400, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != "1 day" {
		t.Errorf("Expected '1 day', got '%s'", result)
	}
}

func TestDuration1Day2Hours(t *testing.T) {
	result, err := Duration(93600, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != "1 day, 2 hours" {
		t.Errorf("Expected '1 day, 2 hours', got '%s'", result)
	}
}

func TestDuration7Days(t *testing.T) {
	result, err := Duration(604800, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != "7 days" {
		t.Errorf("Expected '7 days', got '%s'", result)
	}
}

func TestDuration1Month30Days(t *testing.T) {
	result, err := Duration(2592000, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != "1 month" {
		t.Errorf("Expected '1 month', got '%s'", result)
	}
}

func TestDuration1Year365Days(t *testing.T) {
	result, err := Duration(31536000, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != "1 year" {
		t.Errorf("Expected '1 year', got '%s'", result)
	}
}

func TestDuration1Year2Months(t *testing.T) {
	result, err := Duration(36720000, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != "1 year, 2 months" {
		t.Errorf("Expected '1 year, 2 months', got '%s'", result)
	}
}

func TestDurationCompact1h1m(t *testing.T) {
	result, err := Duration(3661, &DurationOptions{Compact: true})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != "1h 1m" {
		t.Errorf("Expected '1h 1m', got '%s'", result)
	}
}

func TestDurationCompact2h30m(t *testing.T) {
	result, err := Duration(9000, &DurationOptions{Compact: true})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != "2h 30m" {
		t.Errorf("Expected '2h 30m', got '%s'", result)
	}
}

func TestDurationCompact1d2h(t *testing.T) {
	result, err := Duration(93600, &DurationOptions{Compact: true})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != "1d 2h" {
		t.Errorf("Expected '1d 2h', got '%s'", result)
	}
}

func TestDurationCompact45s(t *testing.T) {
	result, err := Duration(45, &DurationOptions{Compact: true})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != "45s" {
		t.Errorf("Expected '45s', got '%s'", result)
	}
}

func TestDurationCompact0s(t *testing.T) {
	result, err := Duration(0, &DurationOptions{Compact: true})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != "0s" {
		t.Errorf("Expected '0s', got '%s'", result)
	}
}

func TestDurationMaxUnits1HoursOnly(t *testing.T) {
	result, err := Duration(3661, &DurationOptions{MaxUnits: 1})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != "1 hour" {
		t.Errorf("Expected '1 hour', got '%s'", result)
	}
}

func TestDurationMaxUnits1DaysOnly(t *testing.T) {
	result, err := Duration(93600, &DurationOptions{MaxUnits: 1})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != "1 day" {
		t.Errorf("Expected '1 day', got '%s'", result)
	}
}

func TestDurationMaxUnits3(t *testing.T) {
	result, err := Duration(93661, &DurationOptions{MaxUnits: 3})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != "1 day, 2 hours, 1 minute" {
		t.Errorf("Expected '1 day, 2 hours, 1 minute', got '%s'", result)
	}
}

func TestDurationCompactMaxUnits1(t *testing.T) {
	result, err := Duration(9000, &DurationOptions{Compact: true, MaxUnits: 1})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != "2h" {
		t.Errorf("Expected '2h', got '%s'", result)
	}
}

func TestDurationErrorNegativeSeconds(t *testing.T) {
	_, err := Duration(-100, nil)
	if err == nil {
		t.Errorf("Expected error for negative seconds")
	}
}

// parse_duration tests
func TestParseDurationCompactHoursMinutes(t *testing.T) {
	result, err := ParseDuration("2h30m")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 9000 {
		t.Errorf("Expected 9000, got %d", result)
	}
}

func TestParseDurationCompactWithSpace(t *testing.T) {
	result, err := ParseDuration("2h 30m")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 9000 {
		t.Errorf("Expected 9000, got %d", result)
	}
}

func TestParseDurationCompactWithComma(t *testing.T) {
	result, err := ParseDuration("2h, 30m")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 9000 {
		t.Errorf("Expected 9000, got %d", result)
	}
}

func TestParseDurationVerbose(t *testing.T) {
	result, err := ParseDuration("2 hours 30 minutes")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 9000 {
		t.Errorf("Expected 9000, got %d", result)
	}
}

func TestParseDurationVerboseWithAnd(t *testing.T) {
	result, err := ParseDuration("2 hours and 30 minutes")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 9000 {
		t.Errorf("Expected 9000, got %d", result)
	}
}

func TestParseDurationVerboseWithCommaAnd(t *testing.T) {
	result, err := ParseDuration("2 hours, and 30 minutes")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 9000 {
		t.Errorf("Expected 9000, got %d", result)
	}
}

func TestParseDurationDecimalHours(t *testing.T) {
	result, err := ParseDuration("2.5 hours")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 9000 {
		t.Errorf("Expected 9000, got %d", result)
	}
}

func TestParseDurationDecimalCompact(t *testing.T) {
	result, err := ParseDuration("1.5h")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 5400 {
		t.Errorf("Expected 5400, got %d", result)
	}
}

func TestParseDurationSingleUnitMinutesVerbose(t *testing.T) {
	result, err := ParseDuration("90 minutes")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 5400 {
		t.Errorf("Expected 5400, got %d", result)
	}
}

func TestParseDurationSingleUnitMinutesCompact(t *testing.T) {
	result, err := ParseDuration("90m")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 5400 {
		t.Errorf("Expected 5400, got %d", result)
	}
}

func TestParseDurationSingleUnitMin(t *testing.T) {
	result, err := ParseDuration("90min")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 5400 {
		t.Errorf("Expected 5400, got %d", result)
	}
}

func TestParseDurationColonNotationHMM(t *testing.T) {
	result, err := ParseDuration("2:30")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 9000 {
		t.Errorf("Expected 9000, got %d", result)
	}
}

func TestParseDurationColonNotationHMMSS(t *testing.T) {
	result, err := ParseDuration("1:30:00")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 5400 {
		t.Errorf("Expected 5400, got %d", result)
	}
}

func TestParseDurationColonNotationWithSeconds(t *testing.T) {
	result, err := ParseDuration("0:05:30")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 330 {
		t.Errorf("Expected 330, got %d", result)
	}
}

func TestParseDurationDaysVerbose(t *testing.T) {
	result, err := ParseDuration("2 days")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 172800 {
		t.Errorf("Expected 172800, got %d", result)
	}
}

func TestParseDurationDaysCompact(t *testing.T) {
	result, err := ParseDuration("2d")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 172800 {
		t.Errorf("Expected 172800, got %d", result)
	}
}

func TestParseDurationWeeksVerbose(t *testing.T) {
	result, err := ParseDuration("1 week")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 604800 {
		t.Errorf("Expected 604800, got %d", result)
	}
}

func TestParseDurationWeeksCompact(t *testing.T) {
	result, err := ParseDuration("1w")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 604800 {
		t.Errorf("Expected 604800, got %d", result)
	}
}

func TestParseDurationMixedVerbose(t *testing.T) {
	result, err := ParseDuration("1 day, 2 hours, and 30 minutes")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 95400 {
		t.Errorf("Expected 95400, got %d", result)
	}
}

func TestParseDurationMixedCompact(t *testing.T) {
	result, err := ParseDuration("1d 2h 30m")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 95400 {
		t.Errorf("Expected 95400, got %d", result)
	}
}

func TestParseDurationSecondsOnlyVerbose(t *testing.T) {
	result, err := ParseDuration("45 seconds")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 45 {
		t.Errorf("Expected 45, got %d", result)
	}
}

func TestParseDurationSecondsCompactS(t *testing.T) {
	result, err := ParseDuration("45s")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 45 {
		t.Errorf("Expected 45, got %d", result)
	}
}

func TestParseDurationSecondsCompactSec(t *testing.T) {
	result, err := ParseDuration("45sec")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 45 {
		t.Errorf("Expected 45, got %d", result)
	}
}

func TestParseDurationHoursHr(t *testing.T) {
	result, err := ParseDuration("2hr")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 7200 {
		t.Errorf("Expected 7200, got %d", result)
	}
}

func TestParseDurationHoursHrs(t *testing.T) {
	result, err := ParseDuration("2hrs")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 7200 {
		t.Errorf("Expected 7200, got %d", result)
	}
}

func TestParseDurationMinutesMins(t *testing.T) {
	result, err := ParseDuration("30mins")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 1800 {
		t.Errorf("Expected 1800, got %d", result)
	}
}

func TestParseDurationCaseInsensitive(t *testing.T) {
	result, err := ParseDuration("2H 30M")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 9000 {
		t.Errorf("Expected 9000, got %d", result)
	}
}

func TestParseDurationWhitespaceTolerance(t *testing.T) {
	result, err := ParseDuration("  2 hours   30 minutes  ")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 9000 {
		t.Errorf("Expected 9000, got %d", result)
	}
}

func TestParseDurationErrorEmptyString(t *testing.T) {
	_, err := ParseDuration("")
	if err == nil {
		t.Errorf("Expected error for empty string")
	}
}

func TestParseDurationErrorNoUnits(t *testing.T) {
	_, err := ParseDuration("hello world")
	if err == nil {
		t.Errorf("Expected error for no units")
	}
}

func TestParseDurationErrorNegative(t *testing.T) {
	_, err := ParseDuration("-5 hours")
	if err == nil {
		t.Errorf("Expected error for negative")
	}
}

func TestParseDurationErrorJustNumber(t *testing.T) {
	_, err := ParseDuration("42")
	if err == nil {
		t.Errorf("Expected error for just number")
	}
}

// human_date tests
func TestHumanDateToday(t *testing.T) {
	result := HumanDate(1705276800, 1705276800)
	if result != "Today" {
		t.Errorf("Expected 'Today', got '%s'", result)
	}
}

func TestHumanDateTodaySameDayDifferentTime(t *testing.T) {
	result := HumanDate(1705320000, 1705276800)
	if result != "Today" {
		t.Errorf("Expected 'Today', got '%s'", result)
	}
}

func TestHumanDateYesterday(t *testing.T) {
	result := HumanDate(1705190400, 1705276800)
	if result != "Yesterday" {
		t.Errorf("Expected 'Yesterday', got '%s'", result)
	}
}

func TestHumanDateTomorrow(t *testing.T) {
	result := HumanDate(1705363200, 1705276800)
	if result != "Tomorrow" {
		t.Errorf("Expected 'Tomorrow', got '%s'", result)
	}
}

func TestHumanDateLastSunday1DayBeforeMonday(t *testing.T) {
	result := HumanDate(1705190400, 1705276800)
	if result != "Yesterday" {
		t.Errorf("Expected 'Yesterday', got '%s'", result)
	}
}

func TestHumanDateLastSaturday2DaysAgo(t *testing.T) {
	result := HumanDate(1705104000, 1705276800)
	if result != "Last Saturday" {
		t.Errorf("Expected 'Last Saturday', got '%s'", result)
	}
}

func TestHumanDateLastFriday3DaysAgo(t *testing.T) {
	result := HumanDate(1705017600, 1705276800)
	if result != "Last Friday" {
		t.Errorf("Expected 'Last Friday', got '%s'", result)
	}
}

func TestHumanDateLastThursday4DaysAgo(t *testing.T) {
	result := HumanDate(1704931200, 1705276800)
	if result != "Last Thursday" {
		t.Errorf("Expected 'Last Thursday', got '%s'", result)
	}
}

func TestHumanDateLastWednesday5DaysAgo(t *testing.T) {
	result := HumanDate(1704844800, 1705276800)
	if result != "Last Wednesday" {
		t.Errorf("Expected 'Last Wednesday', got '%s'", result)
	}
}

func TestHumanDateLastTuesday6DaysAgo(t *testing.T) {
	result := HumanDate(1704758400, 1705276800)
	if result != "Last Tuesday" {
		t.Errorf("Expected 'Last Tuesday', got '%s'", result)
	}
}

func TestHumanDateLastMonday7DaysAgoBecomesDate(t *testing.T) {
	result := HumanDate(1704672000, 1705276800)
	if result != "January 8" {
		t.Errorf("Expected 'January 8', got '%s'", result)
	}
}

func TestHumanDateThisTuesday1DayFuture(t *testing.T) {
	result := HumanDate(1705363200, 1705276800)
	if result != "Tomorrow" {
		t.Errorf("Expected 'Tomorrow', got '%s'", result)
	}
}

func TestHumanDateThisWednesday2DaysFuture(t *testing.T) {
	result := HumanDate(1705449600, 1705276800)
	if result != "This Wednesday" {
		t.Errorf("Expected 'This Wednesday', got '%s'", result)
	}
}

func TestHumanDateThisThursday3DaysFuture(t *testing.T) {
	result := HumanDate(1705536000, 1705276800)
	if result != "This Thursday" {
		t.Errorf("Expected 'This Thursday', got '%s'", result)
	}
}

func TestHumanDateThisSunday6DaysFuture(t *testing.T) {
	result := HumanDate(1705795200, 1705276800)
	if result != "This Sunday" {
		t.Errorf("Expected 'This Sunday', got '%s'", result)
	}
}

func TestHumanDateNextMonday7DaysFutureBecomesDate(t *testing.T) {
	result := HumanDate(1705881600, 1705276800)
	if result != "January 22" {
		t.Errorf("Expected 'January 22', got '%s'", result)
	}
}

func TestHumanDateSameYearDifferentMonth(t *testing.T) {
	result := HumanDate(1709251200, 1705276800)
	if result != "March 1" {
		t.Errorf("Expected 'March 1', got '%s'", result)
	}
}

func TestHumanDateSameYearEndOfYear(t *testing.T) {
	result := HumanDate(1735603200, 1705276800)
	if result != "December 31" {
		t.Errorf("Expected 'December 31', got '%s'", result)
	}
}

func TestHumanDatePreviousYear(t *testing.T) {
	result := HumanDate(1672531200, 1705276800)
	if result != "January 1, 2023" {
		t.Errorf("Expected 'January 1, 2023', got '%s'", result)
	}
}

func TestHumanDateNextYear(t *testing.T) {
	result := HumanDate(1736121600, 1705276800)
	if result != "January 6, 2025" {
		t.Errorf("Expected 'January 6, 2025', got '%s'", result)
	}
}

// date_range tests
func TestDateRangeSameDay(t *testing.T) {
	result := DateRange(1705276800, 1705276800)
	if result != "January 15, 2024" {
		t.Errorf("Expected 'January 15, 2024', got '%s'", result)
	}
}

func TestDateRangeSameDayDifferentTimes(t *testing.T) {
	result := DateRange(1705276800, 1705320000)
	if result != "January 15, 2024" {
		t.Errorf("Expected 'January 15, 2024', got '%s'", result)
	}
}

func TestDateRangeConsecutiveDaysSameMonth(t *testing.T) {
	result := DateRange(1705276800, 1705363200)
	if result != "January 15–16, 2024" {
		t.Errorf("Expected 'January 15–16, 2024', got '%s'", result)
	}
}

func TestDateRangeSameMonthRange(t *testing.T) {
	result := DateRange(1705276800, 1705881600)
	if result != "January 15–22, 2024" {
		t.Errorf("Expected 'January 15–22, 2024', got '%s'", result)
	}
}

func TestDateRangeSameYearDifferentMonths(t *testing.T) {
	result := DateRange(1705276800, 1707955200)
	if result != "January 15 – February 15, 2024" {
		t.Errorf("Expected 'January 15 – February 15, 2024', got '%s'", result)
	}
}

func TestDateRangeDifferentYears(t *testing.T) {
	result := DateRange(1703721600, 1705276800)
	if result != "December 28, 2023 – January 15, 2024" {
		t.Errorf("Expected 'December 28, 2023 – January 15, 2024', got '%s'", result)
	}
}

func TestDateRangeFullYearSpan(t *testing.T) {
	result := DateRange(1704067200, 1735603200)
	if result != "January 1 – December 31, 2024" {
		t.Errorf("Expected 'January 1 – December 31, 2024', got '%s'", result)
	}
}

func TestDateRangeSwappedInputsShouldAutoCorrect(t *testing.T) {
	result := DateRange(1705881600, 1705276800)
	if result != "January 15–22, 2024" {
		t.Errorf("Expected 'January 15–22, 2024', got '%s'", result)
	}
}

func TestDateRangeMultiYearSpan(t *testing.T) {
	result := DateRange(1672531200, 1735689600)
	if result != "January 1, 2023 – January 1, 2025" {
		t.Errorf("Expected 'January 1, 2023 – January 1, 2025', got '%s'", result)
	}
}
