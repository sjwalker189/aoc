package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Could not read puzzle input. %+v\n", err)
		os.Exit(1)
	}

	p1, p2 := Solve(string(input))

	fmt.Println("--- Day 2: Cube Conundrum ---")
	fmt.Printf("Part One: %d\n", p1)
	fmt.Printf("Part Two: %d\n", p2)
}

func Solve(input string) (int, int) {
	sum := 0
	power := 0

	for i, line := range strings.Split(input, "\n") {
		// Game ids can be derived from splice position
		gameId := i + 1

		// Trim the game number prefix as we already know the game id
		input := strings.Split(line, ":")
		if len(input) < 2 {
			continue
		}

		cubes := parseCubeGroups(input[1])

		if isGamePossible(cubes) {
			sum += gameId
		}

		power += powerCubeValue(cubes)

	}

	return sum, power
}

type Cube map[string]int

func NewCube(r, g, b int) Cube {
	return Cube{
		"red":   r,
		"green": g,
		"blue":  b,
	}
}

func parseCubeGroups(input string) []Cube {
	rounds := strings.Split(input, ";")

	results := make([]Cube, len(rounds))

	for i, round := range rounds {
		cube := NewCube(0, 0, 0)

		for _, group := range strings.Split(round, ",") {
			color, count := parseCube(group)
			cube[color] = count

		}

		results[i] = cube
	}

	return results
}

func parseCube(input string) (string, int) {
	pairs := strings.Split(strings.TrimSpace(input), " ")
	count := strings.TrimSpace(pairs[0])
	color := strings.TrimSpace(pairs[1])
	return strings.ToLower(color), mustParseInt(count)
}

func mustParseInt(n string) int {
	r, err := strconv.Atoi(n)
	if err != nil {
		panic(fmt.Errorf("Cannot parse '%s' into int", n))
	}
	return r
}

func isValidCube(c Cube) bool {
	if c["red"] > MAX_RED {
		return false
	}
	if c["green"] > MAX_GREEN {
		return false
	}
	if c["blue"] > MAX_BLUE {
		return false
	}
	return true
}

func isGamePossible(rounds []Cube) bool {
	for _, cube := range rounds {
		if !isValidCube(cube) {
			return false
		}
	}
	return true
}

func powerCubeValue(cubes []Cube) int {
	var r, g, b int

	for _, cube := range cubes {
		if cube["red"] > r {
			r = cube["red"]
		}
		if cube["green"] > g {
			g = cube["green"]
		}
		if cube["blue"] > b {
			b = cube["blue"]
		}
	}

	return r * g * b
}
