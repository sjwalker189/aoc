package main

import (
	"fmt"
	"testing"
)

const RESULT_1 = "142"
const SAMPLE_1 = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

const RESULT_2 = "281"
const SAMPLE_2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func TestPartOne(t *testing.T) {
	result := PartOne(SAMPLE_1)
	fmt.Printf("Part one: %s\n", result)
	if result != RESULT_1 {
		t.Errorf("Expected: %v but received: %v", RESULT_1, result)
	}
}

func TestPartTwo(t *testing.T) {
	result := PartTwo(SAMPLE_2)
	fmt.Printf("Part two: %s\n", result)
	if result != RESULT_2 {
		t.Errorf("Expected %v but received %v", RESULT_2, result)
	}

	// Only one number in the following sequence should be matched
	result2 := PartTwo("oneight")
	fmt.Printf("Part two: %s\n", result2)
	if result2 != "18" {
		t.Errorf("Expected %v but received %v", "18", result2)
	}
}
