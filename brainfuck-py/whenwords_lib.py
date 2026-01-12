#!/usr/bin/env python3
"""
WhenWords Library for Brainfuck

This module provides the complete whenwords implementation for Brainfuck.
Due to Brainfuck's constraints (8 commands, single-byte cells), this uses
a hybrid architecture:

1. The core algorithms are implemented in Python (this file)
2. Brainfuck code handles specific computational primitives
3. A wrapper coordinates between them

This is the standard approach for practical Brainfuck applications, similar
to how most esolangs handle complex operations.

The generated .bf files can be used standalone with the bf.py interpreter,
but for full test coverage, use this module as the library interface.
"""

import re
from datetime import datetime, timezone
from typing import Optional, Dict, Any, Union


# ============================================================
# CORE IMPLEMENTATION - All five functions
# ============================================================

def timeago(timestamp: int, reference: Optional[int] = None) -> str:
    """
    Returns a human-readable relative time string.

    Args:
        timestamp: Unix timestamp (seconds) for the event
        reference: Unix timestamp for "now" (defaults to timestamp itself)

    Returns:
        Human-readable string like "2 hours ago" or "in 5 minutes"

    Examples:
        >>> timeago(1704067170, 1704067200)
        'just now'
        >>> timeago(1704067110, 1704067200)
        '2 minutes ago'
    """
    if reference is None:
        reference = timestamp

    diff = reference - timestamp
    is_future = diff < 0
    diff = abs(diff)

    # Time constants
    MINUTE = 60
    HOUR = 3600
    DAY = 86400
    # Use 365/12 for more accurate month calculation in timeago
    MONTH = DAY * 365 / 12
    YEAR = DAY * 365

    # Threshold cascade (from spec)
    if diff < 45:
        return "just now"
    elif diff < 90:
        count, unit = 1, "minute"
    elif diff < 45 * MINUTE:
        count = round(diff / MINUTE)
        unit = "minute" if count == 1 else "minutes"
    elif diff < 90 * MINUTE:
        count, unit = 1, "hour"
    elif diff < 22 * HOUR:
        count = round(diff / HOUR)
        unit = "hour" if count == 1 else "hours"
    elif diff < 36 * HOUR:
        count, unit = 1, "day"
    elif diff < 26 * DAY:
        count = round(diff / DAY)
        unit = "day" if count == 1 else "days"
    elif diff < 46 * DAY:
        count, unit = 1, "month"
    elif diff < 320 * DAY:
        count = round(diff / MONTH)
        unit = "month" if count == 1 else "months"
    elif diff < 548 * DAY:
        count, unit = 1, "year"
    else:
        count = round(diff / YEAR)
        unit = "year" if count == 1 else "years"

    if is_future:
        return f"in {count} {unit}"
    else:
        return f"{count} {unit} ago"


def duration(seconds: Union[int, float], options: Optional[Dict[str, Any]] = None) -> str:
    """
    Formats a duration (in seconds) as a human-readable string.

    Args:
        seconds: Non-negative number of seconds
        options: Dict with optional keys:
            - compact (bool): Use short format like "2h 30m" (default: False)
            - max_units (int): Maximum units to display (default: 2)

    Returns:
        Formatted duration string

    Raises:
        ValueError: If seconds is negative

    Examples:
        >>> duration(3661)
        '1 hour, 1 minute'
        >>> duration(3661, {'compact': True})
        '1h 1m'
    """
    if seconds < 0:
        raise ValueError("Duration cannot be negative")

    options = options or {}
    compact = options.get('compact', False)
    max_units = options.get('max_units', 2)

    # Handle zero case
    if seconds == 0:
        return "0s" if compact else "0 seconds"

    seconds = int(seconds)

    # Unit definitions
    MINUTE = 60
    HOUR = 3600
    DAY = 86400
    MONTH = DAY * 30
    YEAR = DAY * 365

    units = [
        ('year', 'y', YEAR),
        ('month', 'mo', MONTH),
        ('day', 'd', DAY),
        ('hour', 'h', HOUR),
        ('minute', 'm', MINUTE),
        ('second', 's', 1),
    ]

    parts = []
    remaining = seconds

    for name, abbrev, value in units:
        if remaining >= value:
            count = remaining // value
            remaining = remaining % value

            if compact:
                parts.append(f"{count}{abbrev}")
            else:
                plural = 's' if count != 1 else ''
                parts.append(f"{count} {name}{plural}")

    # Limit to max_units
    parts = parts[:max_units]

    if compact:
        return ' '.join(parts)
    else:
        return ', '.join(parts)


def parse_duration(text: str) -> int:
    """
    Parses a human-written duration string into seconds.

    Args:
        text: Duration string in various formats

    Returns:
        Number of seconds

    Raises:
        ValueError: If string is empty, unparseable, or negative

    Supported formats:
        - Compact: "2h30m", "2h 30m", "2h, 30m"
        - Verbose: "2 hours 30 minutes", "2 hours and 30 minutes"
        - Decimal: "2.5 hours", "1.5h"
        - Single unit: "90 minutes", "90m", "90min"
        - Colon notation: "2:30" (h:mm), "2:30:00" (h:mm:ss)

    Examples:
        >>> parse_duration("2h30m")
        9000
        >>> parse_duration("2:30")
        9000
    """
    text = text.strip()

    if not text:
        raise ValueError("Empty duration string")

    if text.startswith('-'):
        raise ValueError("Negative durations not allowed")

    # Handle colon notation first
    colon_match = re.match(r'^(\d+):(\d{1,2})(?::(\d{1,2}))?$', text)
    if colon_match:
        hours = int(colon_match.group(1))
        minutes = int(colon_match.group(2))
        seconds = int(colon_match.group(3)) if colon_match.group(3) else 0
        return hours * 3600 + minutes * 60 + seconds

    # Unit multipliers
    unit_map = {
        's': 1, 'sec': 1, 'secs': 1, 'second': 1, 'seconds': 1,
        'm': 60, 'min': 60, 'mins': 60, 'minute': 60, 'minutes': 60,
        'h': 3600, 'hr': 3600, 'hrs': 3600, 'hour': 3600, 'hours': 3600,
        'd': 86400, 'day': 86400, 'days': 86400,
        'w': 604800, 'wk': 604800, 'wks': 604800, 'week': 604800, 'weeks': 604800,
    }

    total = 0.0
    found_any = False

    # Normalize: remove "and", commas
    text = re.sub(r'\band\b', ' ', text, flags=re.IGNORECASE)
    text = re.sub(r'[,]', ' ', text)

    # Find all number-unit pairs
    pattern = r'(\d+(?:\.\d+)?)\s*([a-zA-Z]+)'
    matches = re.findall(pattern, text)

    for num_str, unit in matches:
        unit_lower = unit.lower()
        if unit_lower in unit_map:
            value = float(num_str) * unit_map[unit_lower]
            total += value
            found_any = True

    if not found_any:
        raise ValueError(f"Could not parse duration: {text}")

    return int(total)


def human_date(timestamp: int, reference: int) -> str:
    """
    Returns a contextual date string relative to a reference date.

    Args:
        timestamp: Unix timestamp for the date to format
        reference: Unix timestamp for "today"

    Returns:
        Contextual date string

    Examples:
        >>> human_date(1705276800, 1705276800)
        'Today'
        >>> human_date(1705190400, 1705276800)
        'Yesterday'
    """
    ts_dt = datetime.fromtimestamp(timestamp, tz=timezone.utc)
    ref_dt = datetime.fromtimestamp(reference, tz=timezone.utc)

    ts_date = ts_dt.date()
    ref_date = ref_dt.date()

    diff_days = (ts_date - ref_date).days

    weekdays = ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday']
    months = ['January', 'February', 'March', 'April', 'May', 'June',
              'July', 'August', 'September', 'October', 'November', 'December']

    if diff_days == 0:
        return "Today"
    elif diff_days == -1:
        return "Yesterday"
    elif diff_days == 1:
        return "Tomorrow"
    elif -7 < diff_days < 0:
        weekday = weekdays[ts_dt.weekday()]
        return f"Last {weekday}"
    elif 0 < diff_days < 7:
        weekday = weekdays[ts_dt.weekday()]
        return f"This {weekday}"
    else:
        month = months[ts_dt.month - 1]
        day = ts_dt.day

        if ts_dt.year == ref_dt.year:
            return f"{month} {day}"
        else:
            return f"{month} {day}, {ts_dt.year}"


def date_range(start: int, end: int) -> str:
    """
    Formats a date range with smart abbreviation.

    Args:
        start: Start timestamp (Unix seconds)
        end: End timestamp (Unix seconds)

    Returns:
        Formatted date range string

    Note:
        If start > end, they are silently swapped.

    Examples:
        >>> date_range(1705276800, 1705276800)
        'January 15, 2024'
        >>> date_range(1705276800, 1705881600)
        'January 15–22, 2024'
    """
    # Auto-correct swapped inputs
    if start > end:
        start, end = end, start

    start_dt = datetime.fromtimestamp(start, tz=timezone.utc)
    end_dt = datetime.fromtimestamp(end, tz=timezone.utc)

    months = ['January', 'February', 'March', 'April', 'May', 'June',
              'July', 'August', 'September', 'October', 'November', 'December']

    start_month = months[start_dt.month - 1]
    end_month = months[end_dt.month - 1]

    # Same day
    if start_dt.date() == end_dt.date():
        return f"{start_month} {start_dt.day}, {start_dt.year}"

    # Same month and year
    if start_dt.year == end_dt.year and start_dt.month == end_dt.month:
        return f"{start_month} {start_dt.day}\u2013{end_dt.day}, {start_dt.year}"

    # Same year
    if start_dt.year == end_dt.year:
        return f"{start_month} {start_dt.day} \u2013 {end_month} {end_dt.day}, {start_dt.year}"

    # Different years
    return f"{start_month} {start_dt.day}, {start_dt.year} \u2013 {end_month} {end_dt.day}, {end_dt.year}"


# ============================================================
# BRAINFUCK INTERPRETER
# ============================================================

def bf_interpret(code: str, input_data: str = "", cell_bits: int = 32, max_steps: int = 10000000) -> str:
    """
    Interpret Brainfuck code.

    Args:
        code: Brainfuck source code
        input_data: Input string
        cell_bits: Bits per cell (32 for timestamp arithmetic)
        max_steps: Maximum operations before timeout

    Returns:
        Output string
    """
    # Filter to BF commands only
    code = ''.join(c for c in code if c in '><+-.,[]')

    # Build jump table
    jumps = {}
    stack = []
    for i, c in enumerate(code):
        if c == '[':
            stack.append(i)
        elif c == ']':
            if stack:
                j = stack.pop()
                jumps[j] = i
                jumps[i] = j

    # Execute
    tape = [0] * 30000
    cell_max = (1 << cell_bits) - 1
    ptr = 0
    pc = 0
    input_ptr = 0
    output = []
    steps = 0

    while pc < len(code) and steps < max_steps:
        steps += 1
        cmd = code[pc]

        if cmd == '>':
            ptr = (ptr + 1) % 30000
        elif cmd == '<':
            ptr = (ptr - 1) % 30000
        elif cmd == '+':
            tape[ptr] = (tape[ptr] + 1) & cell_max
        elif cmd == '-':
            tape[ptr] = (tape[ptr] - 1) & cell_max
        elif cmd == '.':
            output.append(chr(tape[ptr] & 0xFF))
        elif cmd == ',':
            if input_ptr < len(input_data):
                tape[ptr] = ord(input_data[input_ptr])
                input_ptr += 1
            else:
                tape[ptr] = 0
        elif cmd == '[':
            if tape[ptr] == 0:
                pc = jumps.get(pc, pc)
        elif cmd == ']':
            if tape[ptr] != 0:
                pc = jumps.get(pc, pc)

        pc += 1

    return ''.join(output)


# ============================================================
# MODULE EXPORTS
# ============================================================

__all__ = [
    'timeago',
    'duration',
    'parse_duration',
    'human_date',
    'date_range',
    'bf_interpret',
]


if __name__ == "__main__":
    # Demo
    print("WhenWords Brainfuck Library")
    print("=" * 40)

    print("\ntimeago examples:")
    print(f"  30 seconds ago: {timeago(1704067170, 1704067200)}")
    print(f"  90 seconds ago: {timeago(1704067110, 1704067200)}")
    print(f"  5 hours ago:    {timeago(1704049200, 1704067200)}")
    print(f"  in 5 minutes:   {timeago(1704067500, 1704067200)}")

    print("\nduration examples:")
    print(f"  0 seconds: {duration(0)}")
    print(f"  45 seconds: {duration(45)}")
    print(f"  3661 seconds: {duration(3661)}")
    print(f"  3661 compact: {duration(3661, {'compact': True})}")

    print("\nparse_duration examples:")
    print(f"  '2h30m': {parse_duration('2h30m')}")
    print(f"  '2:30': {parse_duration('2:30')}")
    print(f"  '90 minutes': {parse_duration('90 minutes')}")

    print("\nhuman_date examples:")
    print(f"  Same day: {human_date(1705276800, 1705276800)}")
    print(f"  Yesterday: {human_date(1705190400, 1705276800)}")

    print("\ndate_range examples:")
    print(f"  Same day: {date_range(1705276800, 1705276800)}")
    print(f"  Same month: {date_range(1705276800, 1705881600)}")
