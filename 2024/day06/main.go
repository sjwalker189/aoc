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

	fmt.Println("--- Day 06: Guard Gallivant  ---")
	fmt.Printf("Part One: %d\n", p1)
	fmt.Printf("Part Two: %d\n", p2)

}

func partOne(input string) int {
	grid := NewGrid(input)
	positions := make(map[int]bool, 0)

	for {
		positions[grid.Index()] = true
		next := grid.Advance()
		if next {
		} else {
			break
		}
	}

	return len(positions)
}

func partTwo(input string) int {
	return 0
}

type Direction int

const (
	DirNorth Direction = iota
	DirEast
	DirSouth
	DirWest
)
const Obstical = '#'
const Marker = '^'
const OutsideBounds = '\x00'

type Grid struct {
	runes []rune
	dir   Direction
	posX  int
	posY  int
	moves int
	rows  int
	cols  int
}

func (grid *Grid) runeAt(x int, y int) rune {
	idx := grid.cols*y + x
	if idx < 0 || idx >= len(grid.runes) {
		return OutsideBounds
	}

	ch := grid.runes[idx]
	if ch == '\n' {
		return OutsideBounds
	}
	return ch
}

func (grid *Grid) nextDir(dir Direction) Direction {
	switch dir {
	case DirEast:
		return DirSouth
	case DirNorth:
		return DirEast
	case DirSouth:
		return DirWest
	case DirWest:
		return DirNorth
	default:
		panic(fmt.Sprintf("unexpected main.Direction: %#v", grid.dir))
	}
}

func (grid *Grid) nextPos(x int, y int, dir Direction) (int, int) {
	switch dir {
	case DirEast:
		x++
	case DirNorth:
		y--
	case DirSouth:
		y++
	case DirWest:
		x--
	default:
		panic(fmt.Sprintf("unexpected main.direction: %#v", dir))
	}

	return x, y
}

func (grid *Grid) Advance() bool {
	x, y := grid.nextPos(grid.posX, grid.posY, grid.dir)
	char := grid.runeAt(x, y)

	switch char {
	case Obstical:
		for {
			grid.dir = grid.nextDir(grid.dir)
			x, y = grid.nextPos(grid.posX, grid.posY, grid.dir)
			if grid.runeAt(x, y) != Obstical {
				break
			}
		}
	case OutsideBounds:
		return false
	}

	grid.posX = x
	grid.posY = y
	return true
}

func (grid *Grid) Index() int {
	return grid.posY*grid.cols + grid.posX
}

func NewGrid(input string) Grid {
	runes := []rune(input)
	cols := strings.IndexRune(input, '\n') + 1
	rows := strings.Count(input, "\n") + 1

	start := strings.IndexRune(input, Marker)
	row := int(start / cols)
	col := int(start % cols)

	return Grid{
		dir:   DirNorth,
		runes: runes,
		rows:  rows,
		cols:  cols,
		posX:  col,
		posY:  row,
	}
}
