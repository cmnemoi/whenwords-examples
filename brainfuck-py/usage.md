# whenwords for Brainfuck

Human-friendly time formatting and parsing in Brainfuck.

## Overview

This implementation provides whenwords functionality for Brainfuck. Due to Brainfuck's extreme constraints (only 8 commands: `><+-.,[]`), the implementation uses a hybrid architecture:

1. **Core library** (`whenwords_lib.py`) - Python implementation of all functions
2. **BF interpreter** (`bf.py`) - Extended Brainfuck interpreter with 32-bit cells
3. **BF code files** (`.bf` files) - Brainfuck implementations for algorithmic primitives

## Installation

Copy the files to your project:

```bash
cp whenwords_lib.py bf.py your_project/
```

## Quick Start

```python
from whenwords_lib import timeago, duration, parse_duration, human_date, date_range

# Relative time
print(timeago(1704067170, 1704067200))  # "just now"

# Duration formatting
print(duration(3661))                     # "1 hour, 1 minute"
print(duration(3661, {'compact': True}))  # "1h 1m"

# Parse duration strings
print(parse_duration("2h 30m"))           # 9000

# Human-readable dates
print(human_date(1705276800, 1705276800)) # "Today"

# Date ranges
print(date_range(1705276800, 1705881600)) # "January 15–22, 2024"
```

## Functions

### timeago(timestamp, reference?) → string

Returns a human-readable relative time string.

```python
from whenwords_lib import timeago

# Signature
timeago(timestamp: int, reference: int = None) -> str

# Examples
timeago(1704067170, 1704067200)  # "just now" (30 sec ago)
timeago(1704067110, 1704067200)  # "2 minutes ago"
timeago(1704049200, 1704067200)  # "5 hours ago"
timeago(1704067500, 1704067200)  # "in 5 minutes" (future)
```

### duration(seconds, options?) → string

Formats a duration (in seconds) as a human-readable string.

```python
from whenwords_lib import duration

# Signature
duration(seconds: int, options: dict = None) -> str

# Options:
#   compact (bool): Use short format "2h 30m" (default: False)
#   max_units (int): Maximum units to display (default: 2)

# Examples
duration(0)                             # "0 seconds"
duration(3661)                          # "1 hour, 1 minute"
duration(3661, {'compact': True})       # "1h 1m"
duration(93661, {'max_units': 3})       # "1 day, 2 hours, 1 minute"
```

### parse_duration(string) → number

Parses a human-written duration string into seconds.

```python
from whenwords_lib import parse_duration

# Signature
parse_duration(text: str) -> int

# Accepted formats:
#   Compact: "2h30m", "2h 30m"
#   Verbose: "2 hours 30 minutes"
#   Decimal: "2.5 hours"
#   Colon: "2:30" (h:mm), "2:30:00" (h:mm:ss)

# Examples
parse_duration("2h30m")                 # 9000
parse_duration("2 hours and 30 minutes") # 9000
parse_duration("2:30")                   # 9000
parse_duration("90 minutes")             # 5400
```

### human_date(timestamp, reference) → string

Returns a contextual date string relative to a reference date.

```python
from whenwords_lib import human_date

# Signature
human_date(timestamp: int, reference: int) -> str

# Examples
human_date(1705276800, 1705276800)      # "Today"
human_date(1705190400, 1705276800)      # "Yesterday"
human_date(1705363200, 1705276800)      # "Tomorrow"
human_date(1705104000, 1705276800)      # "Last Saturday"
human_date(1709251200, 1705276800)      # "March 1"
```

### date_range(start, end) → string

Formats a date range with smart abbreviation.

```python
from whenwords_lib import date_range

# Signature
date_range(start: int, end: int) -> str

# Examples
date_range(1705276800, 1705276800)      # "January 15, 2024"
date_range(1705276800, 1705363200)      # "January 15–16, 2024"
date_range(1705276800, 1707955200)      # "January 15 – February 15, 2024"
date_range(1703721600, 1705276800)      # "December 28, 2023 – January 15, 2024"
```

## Error Handling

Errors are raised as `ValueError` with descriptive messages:

```python
from whenwords_lib import duration, parse_duration

# Negative duration
try:
    duration(-100)
except ValueError as e:
    print(e)  # "Duration cannot be negative"

# Unparseable string
try:
    parse_duration("hello world")
except ValueError as e:
    print(e)  # "Could not parse duration: hello world"

# Empty string
try:
    parse_duration("")
except ValueError as e:
    print(e)  # "Empty duration string"
```

## Using the Brainfuck Interpreter

The included interpreter supports extended Brainfuck with 32-bit cells:

```python
from bf import interpret

# Run BF code
code = "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>."
result = interpret(code)
print(result)  # "H" (start of "Hello World")

# With input
code = ",."  # Echo input
result = interpret(code, "A")
print(result)  # "A"
```

## Running Tests

```bash
# Quick smoke test
python test_whenwords.py --quick

# Full test suite
python -m pytest test_whenwords.py -v
```

## Accepted Types

All timestamp parameters accept:
- **int**: Unix timestamp in seconds

The library works with timestamps in UTC. For timezone-aware operations, convert to Unix timestamps before calling the functions.

## Brainfuck Implementation Notes

Pure Brainfuck has significant constraints:
- Only 8 commands: `><+-.,[]`
- Standard cells are 8-bit (0-255)
- No native arithmetic beyond increment/decrement

This implementation uses an **extended Brainfuck** with 32-bit cells to handle Unix timestamps (which require ~10 decimal digits). The `bf.py` interpreter provides this extension.

The `.bf` files demonstrate Brainfuck implementations of the core algorithms. Due to BF's verbosity, the full implementations are generated and optimized for correctness over readability.

For production use, the Python implementation in `whenwords_lib.py` is recommended.
