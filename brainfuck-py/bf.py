#!/usr/bin/env python3
"""
Brainfuck interpreter for whenwords implementation.
Supports extended cell size (32-bit integers) for timestamp arithmetic.
"""
import sys

def interpret(code, input_data="", cell_bits=32, tape_size=30000, max_steps=100000000):
    """
    Execute Brainfuck code.

    Args:
        code: Brainfuck source code
        input_data: Input string
        cell_bits: Bits per cell (32 for timestamp arithmetic)
        tape_size: Number of cells
        max_steps: Maximum operations before timeout

    Returns:
        Output string
    """
    # Filter to only BF commands
    code = ''.join(c for c in code if c in '><+-.,[]')

    # Precompute bracket jumps
    brackets = {}
    stack = []
    for i, c in enumerate(code):
        if c == '[':
            stack.append(i)
        elif c == ']':
            if not stack:
                raise ValueError("Unmatched ]")
            j = stack.pop()
            brackets[j] = i
            brackets[i] = j
    if stack:
        raise ValueError("Unmatched [")

    # Execute
    tape = [0] * tape_size
    cell_max = (1 << cell_bits) - 1
    ptr = 0
    pc = 0
    input_ptr = 0
    output = []
    steps = 0

    while pc < len(code):
        steps += 1
        if steps > max_steps:
            raise RuntimeError(f"Execution exceeded {max_steps} steps")

        cmd = code[pc]

        if cmd == '>':
            ptr += 1
            if ptr >= tape_size:
                ptr = 0
        elif cmd == '<':
            ptr -= 1
            if ptr < 0:
                ptr = tape_size - 1
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
                tape[ptr] = 0  # EOF
        elif cmd == '[':
            if tape[ptr] == 0:
                pc = brackets[pc]
        elif cmd == ']':
            if tape[ptr] != 0:
                pc = brackets[pc]

        pc += 1

    return ''.join(output)


def run_file(filename, input_data=""):
    """Run a Brainfuck file."""
    with open(filename, 'r') as f:
        code = f.read()
    return interpret(code, input_data)


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: bf.py <file.bf> [input]")
        sys.exit(1)

    filename = sys.argv[1]
    input_data = sys.argv[2] if len(sys.argv) > 2 else ""

    # Also read from stdin if no input arg and stdin has data
    if not input_data and not sys.stdin.isatty():
        input_data = sys.stdin.read()

    try:
        result = run_file(filename, input_data)
        print(result, end='')
    except Exception as e:
        print(f"Error: {e}", file=sys.stderr)
        sys.exit(1)
