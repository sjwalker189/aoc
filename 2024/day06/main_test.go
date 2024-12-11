package main

import (
	"fmt"
	"testing"
)

const SAMPLE = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

const RESULT_1 = 41
const RESULT_2 = 1

func TestPartOne(t *testing.T) {
	p1 := partOne(SAMPLE)
	fmt.Printf("Part one: %d\n", p1)
	if p1 != RESULT_1 {
		t.Errorf("Expected: %v but received: %v", RESULT_1, p1)
	}
}
