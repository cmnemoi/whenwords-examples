#!/usr/bin/env python3
"""
Test suite for WhenWords Brainfuck implementation.

This test file is generated from tests.yaml and verifies that the
Brainfuck implementation passes all specification test cases.

Usage:
    python test_whenwords.py           # Run all tests
    python test_whenwords.py -v        # Verbose output
    python test_whenwords.py --quick   # Quick smoke test
"""

import sys
import os

# Add the current directory to path
sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))

from whenwords_lib import timeago, duration, parse_duration, human_date, date_range
import pytest


# ============================================================
# TIMEAGO TESTS (from tests.yaml)
# ============================================================

class TestTimeago:
    """Tests for timeago function."""

    def test_just_now_identical_timestamps(self):
        """just now - identical timestamps"""
        result = timeago(1704067200, 1704067200)
        assert result == "just now"

    def test_just_now_30_seconds_ago(self):
        """just now - 30 seconds ago"""
        result = timeago(1704067170, 1704067200)
        assert result == "just now"

    def test_just_now_44_seconds_ago(self):
        """just now - 44 seconds ago"""
        result = timeago(1704067156, 1704067200)
        assert result == "just now"

    def test_1_minute_ago_45_seconds(self):
        """1 minute ago - 45 seconds"""
        result = timeago(1704067155, 1704067200)
        assert result == "1 minute ago"

    def test_1_minute_ago_89_seconds(self):
        """1 minute ago - 89 seconds"""
        result = timeago(1704067111, 1704067200)
        assert result == "1 minute ago"

    def test_2_minutes_ago_90_seconds(self):
        """2 minutes ago - 90 seconds"""
        result = timeago(1704067110, 1704067200)
        assert result == "2 minutes ago"

    def test_30_minutes_ago(self):
        """30 minutes ago"""
        result = timeago(1704065400, 1704067200)
        assert result == "30 minutes ago"

    def test_44_minutes_ago(self):
        """44 minutes ago"""
        result = timeago(1704064560, 1704067200)
        assert result == "44 minutes ago"

    def test_1_hour_ago_45_minutes(self):
        """1 hour ago - 45 minutes"""
        result = timeago(1704064500, 1704067200)
        assert result == "1 hour ago"

    def test_1_hour_ago_89_minutes(self):
        """1 hour ago - 89 minutes"""
        result = timeago(1704061860, 1704067200)
        assert result == "1 hour ago"

    def test_2_hours_ago_90_minutes(self):
        """2 hours ago - 90 minutes"""
        result = timeago(1704061800, 1704067200)
        assert result == "2 hours ago"

    def test_5_hours_ago(self):
        """5 hours ago"""
        result = timeago(1704049200, 1704067200)
        assert result == "5 hours ago"

    def test_21_hours_ago(self):
        """21 hours ago"""
        result = timeago(1703991600, 1704067200)
        assert result == "21 hours ago"

    def test_1_day_ago_22_hours(self):
        """1 day ago - 22 hours"""
        result = timeago(1703988000, 1704067200)
        assert result == "1 day ago"

    def test_1_day_ago_35_hours(self):
        """1 day ago - 35 hours"""
        result = timeago(1703941200, 1704067200)
        assert result == "1 day ago"

    def test_2_days_ago_36_hours(self):
        """2 days ago - 36 hours"""
        result = timeago(1703937600, 1704067200)
        assert result == "2 days ago"

    def test_7_days_ago(self):
        """7 days ago"""
        result = timeago(1703462400, 1704067200)
        assert result == "7 days ago"

    def test_25_days_ago(self):
        """25 days ago"""
        result = timeago(1701907200, 1704067200)
        assert result == "25 days ago"

    def test_1_month_ago_26_days(self):
        """1 month ago - 26 days"""
        result = timeago(1701820800, 1704067200)
        assert result == "1 month ago"

    def test_1_month_ago_45_days(self):
        """1 month ago - 45 days"""
        result = timeago(1700179200, 1704067200)
        assert result == "1 month ago"

    def test_2_months_ago_46_days(self):
        """2 months ago - 46 days"""
        result = timeago(1700092800, 1704067200)
        assert result == "2 months ago"

    def test_6_months_ago(self):
        """6 months ago"""
        result = timeago(1688169600, 1704067200)
        assert result == "6 months ago"

    def test_10_months_ago_319_days(self):
        """10 months ago - 319 days"""
        result = timeago(1676505600, 1704067200)
        assert result == "10 months ago"

    def test_1_year_ago_320_days(self):
        """1 year ago - 320 days"""
        result = timeago(1676419200, 1704067200)
        assert result == "1 year ago"

    def test_1_year_ago_547_days(self):
        """1 year ago - 547 days"""
        result = timeago(1656806400, 1704067200)
        assert result == "1 year ago"

    def test_2_years_ago_548_days(self):
        """2 years ago - 548 days"""
        result = timeago(1656720000, 1704067200)
        assert result == "2 years ago"

    def test_5_years_ago(self):
        """5 years ago"""
        result = timeago(1546300800, 1704067200)
        assert result == "5 years ago"

    def test_future_in_just_now_30_seconds(self):
        """future - in just now (30 seconds)"""
        result = timeago(1704067230, 1704067200)
        assert result == "just now"

    def test_future_in_1_minute(self):
        """future - in 1 minute"""
        result = timeago(1704067260, 1704067200)
        assert result == "in 1 minute"

    def test_future_in_5_minutes(self):
        """future - in 5 minutes"""
        result = timeago(1704067500, 1704067200)
        assert result == "in 5 minutes"

    def test_future_in_1_hour(self):
        """future - in 1 hour"""
        result = timeago(1704070200, 1704067200)
        assert result == "in 1 hour"

    def test_future_in_3_hours(self):
        """future - in 3 hours"""
        result = timeago(1704078000, 1704067200)
        assert result == "in 3 hours"

    def test_future_in_1_day(self):
        """future - in 1 day"""
        result = timeago(1704150000, 1704067200)
        assert result == "in 1 day"

    def test_future_in_2_days(self):
        """future - in 2 days"""
        result = timeago(1704240000, 1704067200)
        assert result == "in 2 days"

    def test_future_in_1_month(self):
        """future - in 1 month"""
        result = timeago(1706745600, 1704067200)
        assert result == "in 1 month"

    def test_future_in_1_year(self):
        """future - in 1 year"""
        result = timeago(1735689600, 1704067200)
        assert result == "in 1 year"


# ============================================================
# DURATION TESTS (from tests.yaml)
# ============================================================

class TestDuration:
    """Tests for duration function."""

    def test_zero_seconds(self):
        """zero seconds"""
        result = duration(0)
        assert result == "0 seconds"

    def test_1_second(self):
        """1 second"""
        result = duration(1)
        assert result == "1 second"

    def test_45_seconds(self):
        """45 seconds"""
        result = duration(45)
        assert result == "45 seconds"

    def test_1_minute(self):
        """1 minute"""
        result = duration(60)
        assert result == "1 minute"

    def test_1_minute_30_seconds(self):
        """1 minute 30 seconds"""
        result = duration(90)
        assert result == "1 minute, 30 seconds"

    def test_2_minutes(self):
        """2 minutes"""
        result = duration(120)
        assert result == "2 minutes"

    def test_1_hour(self):
        """1 hour"""
        result = duration(3600)
        assert result == "1 hour"

    def test_1_hour_1_minute(self):
        """1 hour 1 minute"""
        result = duration(3661)
        assert result == "1 hour, 1 minute"

    def test_1_hour_30_minutes(self):
        """1 hour 30 minutes"""
        result = duration(5400)
        assert result == "1 hour, 30 minutes"

    def test_2_hours_30_minutes(self):
        """2 hours 30 minutes"""
        result = duration(9000)
        assert result == "2 hours, 30 minutes"

    def test_1_day(self):
        """1 day"""
        result = duration(86400)
        assert result == "1 day"

    def test_1_day_2_hours(self):
        """1 day 2 hours"""
        result = duration(93600)
        assert result == "1 day, 2 hours"

    def test_7_days(self):
        """7 days"""
        result = duration(604800)
        assert result == "7 days"

    def test_1_month_30_days(self):
        """1 month (30 days)"""
        result = duration(2592000)
        assert result == "1 month"

    def test_1_year_365_days(self):
        """1 year (365 days)"""
        result = duration(31536000)
        assert result == "1 year"

    def test_1_year_2_months(self):
        """1 year 2 months"""
        result = duration(36720000)
        assert result == "1 year, 2 months"

    def test_compact_1h_1m(self):
        """compact - 1h 1m"""
        result = duration(3661, {'compact': True})
        assert result == "1h 1m"

    def test_compact_2h_30m(self):
        """compact - 2h 30m"""
        result = duration(9000, {'compact': True})
        assert result == "2h 30m"

    def test_compact_1d_2h(self):
        """compact - 1d 2h"""
        result = duration(93600, {'compact': True})
        assert result == "1d 2h"

    def test_compact_45s(self):
        """compact - 45s"""
        result = duration(45, {'compact': True})
        assert result == "45s"

    def test_compact_0s(self):
        """compact - 0s"""
        result = duration(0, {'compact': True})
        assert result == "0s"

    def test_max_units_1_hours_only(self):
        """max_units 1 - hours only"""
        result = duration(3661, {'max_units': 1})
        assert result == "1 hour"

    def test_max_units_1_days_only(self):
        """max_units 1 - days only"""
        result = duration(93600, {'max_units': 1})
        assert result == "1 day"

    def test_max_units_3(self):
        """max_units 3"""
        result = duration(93661, {'max_units': 3})
        assert result == "1 day, 2 hours, 1 minute"

    def test_compact_max_units_1(self):
        """compact max_units 1"""
        result = duration(9000, {'compact': True, 'max_units': 1})
        assert result == "2h"

    def test_error_negative_seconds(self):
        """error - negative seconds"""
        with pytest.raises(ValueError):
            duration(-100)


# ============================================================
# PARSE_DURATION TESTS (from tests.yaml)
# ============================================================

class TestParseDuration:
    """Tests for parse_duration function."""

    def test_compact_hours_minutes(self):
        """compact hours minutes"""
        result = parse_duration("2h30m")
        assert result == 9000

    def test_compact_with_space(self):
        """compact with space"""
        result = parse_duration("2h 30m")
        assert result == 9000

    def test_compact_with_comma(self):
        """compact with comma"""
        result = parse_duration("2h, 30m")
        assert result == 9000

    def test_verbose(self):
        """verbose"""
        result = parse_duration("2 hours 30 minutes")
        assert result == 9000

    def test_verbose_with_and(self):
        """verbose with and"""
        result = parse_duration("2 hours and 30 minutes")
        assert result == 9000

    def test_verbose_with_comma_and(self):
        """verbose with comma and"""
        result = parse_duration("2 hours, and 30 minutes")
        assert result == 9000

    def test_decimal_hours(self):
        """decimal hours"""
        result = parse_duration("2.5 hours")
        assert result == 9000

    def test_decimal_compact(self):
        """decimal compact"""
        result = parse_duration("1.5h")
        assert result == 5400

    def test_single_unit_minutes_verbose(self):
        """single unit minutes verbose"""
        result = parse_duration("90 minutes")
        assert result == 5400

    def test_single_unit_minutes_compact(self):
        """single unit minutes compact"""
        result = parse_duration("90m")
        assert result == 5400

    def test_single_unit_min(self):
        """single unit min"""
        result = parse_duration("90min")
        assert result == 5400

    def test_colon_notation_h_mm(self):
        """colon notation h:mm"""
        result = parse_duration("2:30")
        assert result == 9000

    def test_colon_notation_h_mm_ss(self):
        """colon notation h:mm:ss"""
        result = parse_duration("1:30:00")
        assert result == 5400

    def test_colon_notation_with_seconds(self):
        """colon notation with seconds"""
        result = parse_duration("0:05:30")
        assert result == 330

    def test_days_verbose(self):
        """days verbose"""
        result = parse_duration("2 days")
        assert result == 172800

    def test_days_compact(self):
        """days compact"""
        result = parse_duration("2d")
        assert result == 172800

    def test_weeks_verbose(self):
        """weeks verbose"""
        result = parse_duration("1 week")
        assert result == 604800

    def test_weeks_compact(self):
        """weeks compact"""
        result = parse_duration("1w")
        assert result == 604800

    def test_mixed_verbose(self):
        """mixed verbose"""
        result = parse_duration("1 day, 2 hours, and 30 minutes")
        assert result == 95400

    def test_mixed_compact(self):
        """mixed compact"""
        result = parse_duration("1d 2h 30m")
        assert result == 95400

    def test_seconds_only_verbose(self):
        """seconds only verbose"""
        result = parse_duration("45 seconds")
        assert result == 45

    def test_seconds_compact_s(self):
        """seconds compact s"""
        result = parse_duration("45s")
        assert result == 45

    def test_seconds_compact_sec(self):
        """seconds compact sec"""
        result = parse_duration("45sec")
        assert result == 45

    def test_hours_hr(self):
        """hours hr"""
        result = parse_duration("2hr")
        assert result == 7200

    def test_hours_hrs(self):
        """hours hrs"""
        result = parse_duration("2hrs")
        assert result == 7200

    def test_minutes_mins(self):
        """minutes mins"""
        result = parse_duration("30mins")
        assert result == 1800

    def test_case_insensitive(self):
        """case insensitive"""
        result = parse_duration("2H 30M")
        assert result == 9000

    def test_whitespace_tolerance(self):
        """whitespace tolerance"""
        result = parse_duration("  2 hours   30 minutes  ")
        assert result == 9000

    def test_error_empty_string(self):
        """error - empty string"""
        with pytest.raises(ValueError):
            parse_duration("")

    def test_error_no_units(self):
        """error - no units"""
        with pytest.raises(ValueError):
            parse_duration("hello world")

    def test_error_negative(self):
        """error - negative"""
        with pytest.raises(ValueError):
            parse_duration("-5 hours")

    def test_error_just_number(self):
        """error - just number"""
        with pytest.raises(ValueError):
            parse_duration("42")


# ============================================================
# HUMAN_DATE TESTS (from tests.yaml)
# ============================================================

class TestHumanDate:
    """Tests for human_date function."""

    def test_today(self):
        """today"""
        result = human_date(1705276800, 1705276800)
        assert result == "Today"

    def test_today_same_day_different_time(self):
        """today - same day different time"""
        result = human_date(1705320000, 1705276800)
        assert result == "Today"

    def test_yesterday(self):
        """yesterday"""
        result = human_date(1705190400, 1705276800)
        assert result == "Yesterday"

    def test_tomorrow(self):
        """tomorrow"""
        result = human_date(1705363200, 1705276800)
        assert result == "Tomorrow"

    def test_last_sunday_1_day_before_monday(self):
        """last Sunday (1 day before Monday)"""
        result = human_date(1705190400, 1705276800)
        assert result == "Yesterday"

    def test_last_saturday_2_days_ago(self):
        """last Saturday (2 days ago)"""
        result = human_date(1705104000, 1705276800)
        assert result == "Last Saturday"

    def test_last_friday_3_days_ago(self):
        """last Friday (3 days ago)"""
        result = human_date(1705017600, 1705276800)
        assert result == "Last Friday"

    def test_last_thursday_4_days_ago(self):
        """last Thursday (4 days ago)"""
        result = human_date(1704931200, 1705276800)
        assert result == "Last Thursday"

    def test_last_wednesday_5_days_ago(self):
        """last Wednesday (5 days ago)"""
        result = human_date(1704844800, 1705276800)
        assert result == "Last Wednesday"

    def test_last_tuesday_6_days_ago(self):
        """last Tuesday (6 days ago)"""
        result = human_date(1704758400, 1705276800)
        assert result == "Last Tuesday"

    def test_last_monday_7_days_ago_becomes_date(self):
        """last Monday (7 days ago) - becomes date"""
        result = human_date(1704672000, 1705276800)
        assert result == "January 8"

    def test_this_tuesday_1_day_future(self):
        """this Tuesday (1 day future)"""
        result = human_date(1705363200, 1705276800)
        assert result == "Tomorrow"

    def test_this_wednesday_2_days_future(self):
        """this Wednesday (2 days future)"""
        result = human_date(1705449600, 1705276800)
        assert result == "This Wednesday"

    def test_this_thursday_3_days_future(self):
        """this Thursday (3 days future)"""
        result = human_date(1705536000, 1705276800)
        assert result == "This Thursday"

    def test_this_sunday_6_days_future(self):
        """this Sunday (6 days future)"""
        result = human_date(1705795200, 1705276800)
        assert result == "This Sunday"

    def test_next_monday_7_days_future_becomes_date(self):
        """next Monday (7 days future) - becomes date"""
        result = human_date(1705881600, 1705276800)
        assert result == "January 22"

    def test_same_year_different_month(self):
        """same year different month"""
        result = human_date(1709251200, 1705276800)
        assert result == "March 1"

    def test_same_year_end_of_year(self):
        """same year end of year"""
        result = human_date(1735603200, 1705276800)
        assert result == "December 31"

    def test_previous_year(self):
        """previous year"""
        result = human_date(1672531200, 1705276800)
        assert result == "January 1, 2023"

    def test_next_year(self):
        """next year"""
        result = human_date(1736121600, 1705276800)
        assert result == "January 6, 2025"


# ============================================================
# DATE_RANGE TESTS (from tests.yaml)
# ============================================================

class TestDateRange:
    """Tests for date_range function."""

    def test_same_day(self):
        """same day"""
        result = date_range(1705276800, 1705276800)
        assert result == "January 15, 2024"

    def test_same_day_different_times(self):
        """same day different times"""
        result = date_range(1705276800, 1705320000)
        assert result == "January 15, 2024"

    def test_consecutive_days_same_month(self):
        """consecutive days same month"""
        result = date_range(1705276800, 1705363200)
        assert result == "January 15\u201316, 2024"

    def test_same_month_range(self):
        """same month range"""
        result = date_range(1705276800, 1705881600)
        assert result == "January 15\u201322, 2024"

    def test_same_year_different_months(self):
        """same year different months"""
        result = date_range(1705276800, 1707955200)
        assert result == "January 15 \u2013 February 15, 2024"

    def test_different_years(self):
        """different years"""
        result = date_range(1703721600, 1705276800)
        assert result == "December 28, 2023 \u2013 January 15, 2024"

    def test_full_year_span(self):
        """full year span"""
        result = date_range(1704067200, 1735603200)
        assert result == "January 1 \u2013 December 31, 2024"

    def test_swapped_inputs_should_auto_correct(self):
        """swapped inputs - should auto-correct"""
        result = date_range(1705881600, 1705276800)
        assert result == "January 15\u201322, 2024"

    def test_multi_year_span(self):
        """multi-year span"""
        result = date_range(1672531200, 1735689600)
        assert result == "January 1, 2023 \u2013 January 1, 2025"


# ============================================================
# MAIN
# ============================================================

def run_quick_test():
    """Run a quick smoke test."""
    print("Running quick smoke test...")

    tests = [
        ("timeago(1704067170, 1704067200)", timeago(1704067170, 1704067200), "just now"),
        ("duration(3661)", duration(3661), "1 hour, 1 minute"),
        ("parse_duration('2h30m')", parse_duration("2h30m"), 9000),
        ("human_date(1705276800, 1705276800)", human_date(1705276800, 1705276800), "Today"),
        ("date_range(1705276800, 1705881600)", date_range(1705276800, 1705881600), "January 15\u201322, 2024"),
    ]

    passed = 0
    failed = 0

    for name, result, expected in tests:
        if result == expected:
            print(f"  PASS: {name}")
            passed += 1
        else:
            print(f"  FAIL: {name}")
            print(f"    Expected: {expected!r}")
            print(f"    Got:      {result!r}")
            failed += 1

    print(f"\nResults: {passed} passed, {failed} failed")
    return failed == 0


if __name__ == "__main__":
    if "--quick" in sys.argv:
        success = run_quick_test()
        sys.exit(0 if success else 1)
    else:
        # Run pytest
        sys.exit(pytest.main([__file__, "-v"] + sys.argv[1:]))
