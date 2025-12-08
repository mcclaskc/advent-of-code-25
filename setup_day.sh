#!/bin/bash

# Script to set up a new Advent of Code day
# Usage: ./setup_day.sh <day_number>
# Example: ./setup_day.sh 6

if [ $# -eq 0 ]; then
    echo "Usage: $0 <day_number>"
    echo "Example: $0 6"
    exit 1
fi

DAY=$1
DAY_DIR="cmd/day${DAY}"

# Check if day already exists
if [ -d "$DAY_DIR" ]; then
    echo "Error: $DAY_DIR already exists!"
    exit 1
fi

# Create directory structure
mkdir -p "${DAY_DIR}/part1"
mkdir -p "${DAY_DIR}/part2"

# Create empty input.txt
touch "${DAY_DIR}/input.txt"
touch "${DAY_DIR}/test_input.txt"

# Create empty main.go files
touch "${DAY_DIR}/part1/main.go"
touch "${DAY_DIR}/part2/main.go"

# Create empty main_test.go files
touch "${DAY_DIR}/part1/main_test.go"
touch "${DAY_DIR}/part2/main_test.go"

echo "Successfully created structure for day ${DAY}:"
echo "  ${DAY_DIR}/"
echo "    input.txt"
echo "    part1/"
echo "      main.go"
echo "      main_test.go"
echo "    part2/"
echo "      main.go"
echo "      main_test.go"

