package main

import (
	"fmt"
	"maps"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Could not read puzzle input. %+v\n", err)
		os.Exit(1)
	}

	p1, p2 := Solve(string(input))

	fmt.Println("--- Day 1: Historian Hysteria ---")
	fmt.Printf("Part One: %d\n", p1)
	fmt.Printf("Part Two: %d\n", p2)

}

func Solve(input string) (int, int) {
	lhs, rhs := mustParseCoordinates(input)
	distance := calculateDistance(lhs, rhs)
	similarity := calculateSimilarity(lhs, rhs)

	return distance, similarity
}

func mustParseCoordinates(input string) ([]int, []int) {
	lhs := make([]int, 0)
	rhs := make([]int, 0)

	for _, line := range strings.Split(input, "\n") {
		values := strings.Fields(line)
		if len(values) == 0 {
			continue
		}
		lhs = append(lhs, mustParseInt(values[0]))
		rhs = append(rhs, mustParseInt(values[1]))
	}

	sort.Ints(lhs)
	sort.Ints(rhs)

	return lhs, rhs
}

func mustParseInt(n string) int {
	value, err := strconv.Atoi(n)
	if err != nil {
		panic(fmt.Errorf("%s is not a number", n))
	}
	return value
}

func calculateDistance(lhs, rhs []int) int {
	sum := 0
	for i, a := range lhs {
		b := rhs[i]

		// math.Abs() returns float64 so unsure the ideom here for other number types
		if a < b {
			sum += b - a
		} else {
			sum += a - b
		}
	}
	return sum
}

func calculateSimilarity(lhs, rhs []int) int {
	multipliers := make(map[int]int)
	occurences := make(map[int]int)

	// Reduce the left-hand coordinates into a map of times seen in the right-hand
	// coordinates. Because we're using a map here we cannot have duplicates.
	// We only need to scan rhs once, so keep track of duplicates in lhs and use
	// this as a multiplier for the final value
	for _, n := range lhs {
		occurences[n] = 0

		// Increase the multiplier for this number if we have seen it previously
		value, ok := multipliers[n]
		if ok {
			multipliers[n] = value + 1
		} else {
			multipliers[n] = 1
		}
	}

	//
	for _, n := range rhs {
		_, ok := occurences[n]
		if ok {
			occurences[n] += 1
		}
	}

	sum := 0
	for key, value := range maps.All(occurences) {
		// It's impossible for key not to be in multipliers
		sum += ((key * value) * multipliers[key])
	}

	return sum
}
