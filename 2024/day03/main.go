package main

import (
	"aoc/lexer"
	"aoc/parser"
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Could not read puzzle input. %+v\n", err)
		os.Exit(1)
	}

	p1 := partOne(string(input))
	p2 := partTwo(string(input))

	fmt.Println("--- Day 03: Mull It Over ---")
	fmt.Printf("Part One: %d\n", p1)
	fmt.Printf("Part Two: %d\n", p2)
}

func parse(input string) []parser.Operator {
	lex := lexer.NewLexer(input)
	tokens := lex.Tokens()
	parser := parser.NewParser(tokens)
	return parser.Parse()
}

func partOne(input string) int {
	sum := 0
	for _, op := range parse(input) {
		sum += op.Result()
	}
	return sum
}

func partTwo(input string) int {
	sum := 0
	ignore := false

	for _, op := range parse(input) {
		switch op.Name() {
		case parser.OpDo:
			ignore = false
		case parser.OpDont:
			ignore = true
		case parser.OpMultiply:
			if ignore == false {
				sum += op.Result()
			}
		}
	}

	return sum
}
