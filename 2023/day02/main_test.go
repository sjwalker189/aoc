package main

import (
	"fmt"
	"testing"
)

const SAMPLE = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

const RESULT_1 = 8
const RESULT_2 = 2286

func TestDayTwo(t *testing.T) {
	p1, p2 := Solve(SAMPLE)

	fmt.Printf("Part one: %d\n", p1)
	if p1 != RESULT_1 {
		t.Errorf("Expected: %v but received: %v", RESULT_1, p1)
	}

	fmt.Printf("Part two %d\n", p2)
	if p2 != RESULT_2 {
		t.Errorf("Expected: %v but received: %v", RESULT_2, p2)
	}
}
