package parser

import (
	"aoc/lexer"
	"fmt"
	"strconv"
)

func NewParser(tokens []lexer.Token) Parser {
	return Parser{
		tokens: tokens,
		ops:    make([]Operator, 0),
		op:     nil,
	}
}

type Parser struct {
	tokens []lexer.Token
	ops    []Operator
	op     Operator

	expected []lexer.TokenType
}

func (parser *Parser) expectOne(kind lexer.TokenType) {
	parser.expected = []lexer.TokenType{kind}
}

func (parser *Parser) expectAny(kind []lexer.TokenType) {
	parser.expected = kind
}

func (parser *Parser) clear() {
	parser.expectAny([]lexer.TokenType{lexer.IdentMult, lexer.IdentDont, lexer.IdentDo})
	parser.op = nil
}

func (parser *Parser) isExpected(kind lexer.TokenType) bool {
	for _, e := range parser.expected {
		if e == kind {
			return true
		}
	}
	return false
}

func (parser *Parser) Parse() []Operator {
	parser.clear()

	for i, tok := range parser.tokens {
		if !parser.isExpected(tok.Kind) {
			parser.clear()
			continue
		}

		switch tok.Kind {
		case lexer.LParen:
			parser.expectAny([]lexer.TokenType{lexer.Int, lexer.RParen})
		case lexer.Comma:
			parser.expectOne(lexer.Int)
		case lexer.RParen:
			// This parser is really simple, when we reach the end of a function
			// we capture the operation and reset the parser state to look for
			// the next function
			if parser.op != nil {
				parser.ops = append(parser.ops, parser.op)
			}
			parser.clear()
		case lexer.Int:
			arg, err := strconv.Atoi(tok.Raw)
			if err != nil {
				// This should never happen as lexer.go should capture int sequences
				panic(err)
			}

			// We're only handling arguments for lexer.IdentMult so if we've
			// reached this point we need to parse the next sequence of ", int )"
			// Functions can may/must only have two int arguments
			if parser.tokens[i-1].Kind == lexer.LParen {
				parser.op.SetArg1(arg)
				parser.expectOne(lexer.Comma)
			} else {
				parser.op.SetArg2(arg)
				parser.expectOne(lexer.RParen)
			}
		default:
			op, err := NewOperator(tok)
			if err != nil {
				fmt.Println(err)
				parser.clear()
			} else {
				parser.op = op
				parser.expectOne(lexer.LParen)
			}
		}
	}

	return parser.ops
}
