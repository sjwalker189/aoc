package main

import (
	"fmt"
	"testing"
)

const SAMPLE_1 = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
const SAMPLE_2 = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

const RESULT_1 = 161
const RESULT_2 = 48

func TestSolution(t *testing.T) {
	p1 := partOne(SAMPLE_1)
	fmt.Printf("Part one: %d\n", p1)
	if p1 != RESULT_1 {
		t.Errorf("Expected: %v but received: %v", RESULT_1, p1)
	}

	p2 := partTwo(SAMPLE_2)
	fmt.Printf("Part two %d\n", p2)
	if p2 != RESULT_2 {
		t.Errorf("Expected: %v but received: %v", RESULT_2, p2)
	}
}
