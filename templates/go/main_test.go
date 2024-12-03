package main

import (
	"fmt"
	"testing"
)

const SAMPLE = ``

const RESULT_1 = 1
const RESULT_2 = 1

func TestSolution(t *testing.T) {
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
