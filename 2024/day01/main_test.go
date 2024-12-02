package main

import (
	"fmt"
	"testing"
)

const SAMPLE = `3   4
4   3
2   5
1   3
3   9
3   3`

const RESULT_1 = 11
const RESULT_2 = 31

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
