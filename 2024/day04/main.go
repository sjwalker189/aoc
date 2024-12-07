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

	p1 := partOne(string(input))
	p2 := partTwo(string(input))

	fmt.Println("--- Day 04: Ceres Search ---")
	fmt.Printf("Part One: %d\n", p1)
	fmt.Printf("Part Two: %d\n", p2)
}

func partOne(input string) int {
	grid := NewGrid(input)
	word := []rune("XMAS")

	var matches int

	for x := range grid.Cols {
		for y := range grid.Rows {
			matches += grid.Search(x, y, DirNorth, word)
			matches += grid.Search(x, y, DirNorthEast, word)
			matches += grid.Search(x, y, DirEast, word)
			matches += grid.Search(x, y, DirSouthEast, word)
			matches += grid.Search(x, y, DirSouth, word)
			matches += grid.Search(x, y, DirSouthWest, word)
			matches += grid.Search(x, y, DirWest, word)
			matches += grid.Search(x, y, DirNorthWest, word)
		}
	}

	return matches
}

func partTwo(input string) int {
	grid := NewGrid(input)

	// Define each rune as a variable to avoid allocations
	M := 'M'
	A := 'A'
	S := 'S'

	var matches int

	for x := range grid.Cols {
		for y := range grid.Rows {

			// We only need to search when we enounter the middle character
			if grid.RuneAt(x, y) != A {
				continue
			}

			matchLeft := false
			if (grid.RuneAt(x-1, y-1) == M && grid.RuneAt(x+1, y+1) == S) ||
				(grid.RuneAt(x-1, y-1) == S && grid.RuneAt(x+1, y+1) == M) {
				matchLeft = true
			}

			matchRight := false
			if (grid.RuneAt(x+1, y-1) == M && grid.RuneAt(x-1, y+1) == S) ||
				(grid.RuneAt(x+1, y-1) == S && grid.RuneAt(x-1, y+1) == M) {
				matchRight = true
			}

			if matchLeft && matchRight {
				matches++
			}
		}
	}

	return matches
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

func (grid *Grid) Search(x int, y int, dir Direction, word []rune) int {
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
