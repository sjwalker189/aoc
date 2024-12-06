package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Could not read puzzle input. %+v\n", err)
		os.Exit(1)
	}

	p1, p2 := Solve(string(input))

	fmt.Println("--- Day 04: Ceres Search ---")
	fmt.Printf("Part One: %d\n", p1)
	fmt.Printf("Part Two: %d\n", p2)

}

func Solve(input string) (int, int) {
	grid := NewGrid(input)
	return partOne(&grid), 0
}

func partOne(grid *Grid) int {
	return grid.CountWord([]rune("XMAS"))
}

type Direction int

const (
	DirNorth Direction = iota
	DirNorthEast
	DirEast
	DirSouthEast
	DirSouth
	DirSouthWest
	DirWest
	DirNorthWest
)

type Grid struct {
	Runes []rune
	Rows  int
	Cols  int
}

func (grid *Grid) search(x int, y int, dir Direction, word []rune) int {
	for _, char := range word {
		// Ensure we're within the grid bounds
		if x < 0 || y < 0 || y > grid.Rows || x > grid.Cols {
			return 0
		}

		// When not a match exit immediately
		if grid.RuneAt(x, y) != char {
			return 0
		}

		switch dir {
		case DirEast:
			x++
		case DirNorth:
			y--
		case DirNorthEast:
			x++
			y--
		case DirNorthWest:
			x--
			y--
		case DirSouth:
			y++
		case DirSouthEast:
			x++
			y++
		case DirSouthWest:
			x--
			y++
		case DirWest:
			x--
		default:
			panic(fmt.Sprintf("unexpected main.Direction: %#v", dir))
		}
	}

	return 1
}

func (grid *Grid) RuneAt(x int, y int) rune {
	idx := grid.Cols*y + x
	if idx < 0 || idx >= len(grid.Runes) {
		// FIXME: This is a bit hacky but it works for this puzzle.
		// Should return (rune, bool) or similar
		return '\x00'
	}

	return grid.Runes[idx]
}

func (grid *Grid) CountWord(word []rune) int {
	var matches int

	for x := range grid.Cols {
		for y := range grid.Rows {
			// We're returning a 0 or 1 for convenience
			matches += grid.search(x, y, DirNorth, word)
			matches += grid.search(x, y, DirNorthEast, word)
			matches += grid.search(x, y, DirEast, word)
			matches += grid.search(x, y, DirSouthEast, word)
			matches += grid.search(x, y, DirSouth, word)
			matches += grid.search(x, y, DirSouthWest, word)
			matches += grid.search(x, y, DirWest, word)
			matches += grid.search(x, y, DirNorthWest, word)
		}
	}

	return matches
}

func NewGrid(input string) Grid {
	runes := []rune(input)
	cols := strings.IndexRune(input, '\n') + 1
	rows := strings.Count(input, "\n") + 1

	return Grid{
		Runes: runes,
		Rows:  rows,
		Cols:  cols,
	}
}
