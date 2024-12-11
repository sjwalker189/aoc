package main

import (
	"fmt"
	"testing"
)

const SAMPLE = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

const RESULT_1 = 143
const RESULT_2 = 123

func TestSolution(t *testing.T) {
	p1 := partOne(SAMPLE)
	fmt.Printf("Part one: %d\n", p1)
	if p1 != RESULT_1 {
		t.Errorf("Expected: %v but received: %v", RESULT_1, p1)
	}

	p2 := partTwo(SAMPLE)
	fmt.Printf("Part two %d\n", p2)
	if p2 != RESULT_2 {
		t.Errorf("Expected: %v but received: %v", RESULT_2, p2)
	}
}
