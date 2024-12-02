package main

import (
	"fmt"
	"testing"
)

const SAMPLE = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

const RESULT_1 = 2
const RESULT_2 = 4

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
