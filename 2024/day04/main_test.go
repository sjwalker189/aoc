package main

import (
	"fmt"
	"testing"
)

const SAMPLE = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

const RESULT_1 = 18
const RESULT_2 = 9

func TestPartOne(t *testing.T) {
	p1 := partOne(SAMPLE)
	fmt.Printf("Part one: %d\n", p1)
	if p1 != RESULT_1 {
		t.Errorf("Expected: %v but received: %v", RESULT_1, p1)
	}
}

func TestPartTwo(t *testing.T) {
	p2 := partTwo(SAMPLE)
	fmt.Printf("Part two %d\n", p2)
	if p2 != RESULT_2 {
		t.Errorf("Expected: %v but received: %v", RESULT_2, p2)
	}
}
