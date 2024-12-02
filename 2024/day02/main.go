package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const READING_THRESHOLD = 3

type Report = []int

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Could not read puzzle input. %+v\n", err)
		os.Exit(1)
	}

	p1, p2 := Solve(string(input))

	fmt.Println("--- Day 2: Red-Nosed Reports ---")
	fmt.Printf("Part One: %d\n", p1)
	fmt.Printf("Part Two: %d\n", p2)
}

func Solve(input string) (int, int) {
	reports := parseReports(input)

	return len(pickSafeReports(reports)), len(pickSafeReportsWithDampening(reports))
}

func parseReports(input string) []Report {
	reports := make([]Report, 0)

	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}

		levels := make(Report, 0)
		for _, level := range strings.Fields(line) {
			if len(level) == 0 {
				continue
			}
			levels = append(levels, mustParseInt(level))
		}

		reports = append(reports, levels)
	}

	return reports
}

func mustParseInt(n string) int {
	value, err := strconv.Atoi(n)
	if err != nil {
		panic(fmt.Errorf("%s is not a number", n))
	}
	return value
}

func pickSafeReports(reports []Report) []Report {
	safe := make([]Report, 0)

	for _, report := range reports {
		if isSafeReport(report) {
			safe = append(safe, report)
		}
	}

	return safe
}

func pickSafeReportsWithDampening(reports []Report) []Report {
	safe := make([]Report, 0)

	for _, report := range reports {
		// Report is safe, no dampening required
		if isSafeReport(report) {
			safe = append(safe, report)
			continue
		}

		// Drop each recorded level in the report and test until the report is safe
		for i := range len(report) {
			dampened := slices.Delete(slices.Clone(report), i, i+1)
			if isSafeReport(dampened) {
				safe = append(safe, report)
				break
			}
		}
	}

	return safe
}

func isSafeReport(report Report) bool {
	dir := 0

	for i, level := range report {
		// Skip first level as we're comparing from right to left
		if i == 0 {
			continue
		}

		prev := report[i-1]

		// Report is not safe because it is neither increasing or decreasing
		if level == prev {
			return false
		}

		// Determine asc or desc progression
		if dir == 0 {
			if level > prev {
				dir = 1
			} else if level < prev {
				dir = -1
			}
		}

		if dir == 1 {
			// Unsafe because we should be ascending but we encountered a descending value
			if level < prev {
				return false
			}

			if level-prev > READING_THRESHOLD {
				return false
			}
		}

		if dir == -1 {
			// Unsafe because we should be descending but we encountered an ascending value
			if level > prev {
				return false
			}

			if prev-level > READING_THRESHOLD {
				return false
			}
		}
	}

	return true
}
