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

	p1, p2 := Solve(string(input))

	fmt.Println("--- Day 03: Mull It Over ---")
	fmt.Printf("Part One: %d\n", p1)
	fmt.Printf("Part Two: %d\n", p2)
}

func Solve(input string) (int, int) {
	lex := lexer.NewLexer(input)
	tokens := lex.Tokens()
	parser := parser.NewParser(tokens)
	value := parser.Evaluate()

	return value, 0
}
