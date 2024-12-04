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

	expected lexer.TokenType
}

func (parser *Parser) expectNext(kind lexer.TokenType) {
	parser.expected = kind
}

func (parser *Parser) clear() {
	parser.expected = lexer.IdentMult
	parser.op = nil
}

func (parser *Parser) Evaluate() int {
	parser.clear()

	for i, tok := range parser.tokens {
		if tok.Kind != parser.expected {
			parser.clear()
			continue
		}

		switch tok.Kind {
		case lexer.IdentMult:
			op, err := NewOperator(tok)
			if err != nil {
				fmt.Println(err)
				parser.clear()
			} else {
				parser.op = op
				parser.expectNext(lexer.LParen)
			}
		case lexer.LParen:
			parser.expectNext(lexer.Int)
		case lexer.Comma:
			parser.expectNext(lexer.Int)
		case lexer.RParen:
			if parser.op != nil {
				parser.ops = append(parser.ops, parser.op)
				parser.op = nil
			}
			parser.expectNext(lexer.IdentMult)
		case lexer.Int:
			arg, err := strconv.Atoi(tok.Raw)
			if err != nil {
				parser.clear()
				continue
			}
			if parser.tokens[i-1].Kind == lexer.LParen {
				parser.op.SetArg1(arg)
				parser.expectNext(lexer.Comma)
			} else {
				parser.op.SetArg2(arg)
				parser.expectNext(lexer.RParen)
			}
		}
	}

	sum := 0
	for _, op := range parser.ops {
		sum += op.Result()
	}

	return sum
}
