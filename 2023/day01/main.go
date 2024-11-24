package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("--- Day 1: Trebuchet?! ---")
	fmt.Printf("Part One: %s\n", PartOne(puzzleInput()))
	fmt.Printf("Part Two: %s\n", PartTwo(puzzleInput()))
}

func puzzleInput() string {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Could not read puzzle input. %+v\n", err)
		os.Exit(1)
	}
	return string(input)
}

func PartOne(input string) string {
	sum := 0
	for i, line := range strings.Split(input, "\n") {
		numbers := make([]string, 0)
		for _, c := range line {
			if unicode.IsDigit(c) {
				numbers = append(numbers, string(c))
			}
		}

		if len(numbers) == 0 {
			continue
		}

		result, err := strconv.Atoi(numbers[0] + numbers[len(numbers)-1])
		if err != nil {
			fmt.Printf("No matches on line: %d", i)
			continue
		}

		sum += result
	}

	return fmt.Sprintf("%v", sum)
}

func PartTwo(input string) string {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		sum += parseLine(line)
	}
	return fmt.Sprintf("%v", sum)
}

// Match all digit characters and named numbers under 10
var re = regexp.MustCompile("(?:one|two|three|four|five|six|seven|eight|nine|zero|[0-9])")

func parseLine(input string) int {
	matches := make([]string, 0)

	// Itterate the line input to progressively find matches
	// This is needed because named numbers may overlay, for example: "oneight"
	start := 0
	for start < len(input) {
		// Find the next match starting from the current position
		match := re.FindString(input[start:])
		if match == "" {
			break
		}

		// Move the start possition ahead for the next match
		if len(match) == 1 {
			start += len(match)
		} else {
			// When we match a named number move the start position to
			// the last letter. This ensures we can match overlapping numbers
			start += len(match) - 1
		}

		matches = append(matches, match)
	}

	size := len(matches)
	if size == 0 {
		return 0
	}

	lhs := parseDigitString(matches[0])
	rhs := parseDigitString(matches[size-1])

	// Convert the combined string number to an int
	result, err := strconv.Atoi(lhs + rhs)
	if err != nil {
		return 0
	}
	return result
}

var digitMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
	"zero":  "0",
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
	"0":     "0",
}

func parseDigitString(n string) string {
	if digit, ok := digitMap[n]; ok {
		return digit
	}
	panic(fmt.Errorf("'%s' is not a number", n))
}
