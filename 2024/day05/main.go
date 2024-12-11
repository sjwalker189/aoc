package main

import (
	"fmt"
	"os"
	"slices"
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

	p1 := partOne(string(input))
	p2 := partTwo(string(input))

	fmt.Println("--- Day 05: Print Queue ---")
	fmt.Printf("Part One: %d\n", p1)
	fmt.Printf("Part Two: %d\n", p2)
}

type Rule struct {
	Left  int
	Right int
}

type Update []int

func parseSafteyManual(input string) ([]Rule, []Update) {
	parts := strings.Split(input, "\n\n")
	if len(parts) < 2 {
		panic(fmt.Errorf("Invalid input format"))
	}

	ruleset := make([]Rule, 0)
	for _, line := range strings.Split(parts[0], "\n") {
		if len(line) <= 0 {
			continue
		}

		values := strings.Split(line, "|")
		ruleset = append(ruleset, Rule{
			Left:  mustParseInt(values[0]),
			Right: mustParseInt(values[1]),
		})
	}

	updates := make([]Update, 0)
	for _, line := range strings.Split(parts[1], "\n") {
		if len(line) <= 0 {
			continue
		}
		update := make(Update, 0)
		for _, digit := range strings.Split(line, ",") {
			if len(line) <= 0 {
				continue
			}
			update = append(update, mustParseInt(digit))
		}
		updates = append(updates, update)
	}

	return ruleset, updates
}

func mustParseInt(value string) int {
	num, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return num
}

func partOne(input string) int {
	rules, updates := parseSafteyManual(input)

	sum := 0

	for _, update := range updates {
		if isCorrectlyOrdered(rules, update) {
			mid := len(update) / 2
			sum += update[mid]
		}
	}

	return sum
}

func partTwo(input string) int {
	rules, updates := parseSafteyManual(input)

	sum := 0

	for _, update := range updates {
		if !isCorrectlyOrdered(rules, update) {
			// TODO
			// mid := len(update) / 2
			// sum += update[mid]
		}
	}

	return sum
}

func isCorrectlyOrdered(ruleset []Rule, update Update) bool {
	for _, rule := range ruleset {
		left := slices.Index(update, rule.Left)
		right := slices.Index(update, rule.Right)

		if left == -1 || right == -1 {
			continue
		}

		if left < right {
			continue
		}

		return false
	}

	return true
}
