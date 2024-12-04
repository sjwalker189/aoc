package parser

import (
	"aoc/lexer"
	"errors"
)

type Operator interface {
	Result() int
	SetArg1(n int)
	SetArg2(n int)
}

func NewOperator(token lexer.Token) (Operator, error) {
	if token.Kind == lexer.IdentMult {
		return &MultiplicationOperator{}, nil
	}
	return nil, errors.New("Invalid operator type")
}

type MultiplicationOperator struct {
	Arg1 int
	Arg2 int
}

func (op *MultiplicationOperator) SetArg1(n int) {
	op.Arg1 = n
}

func (op *MultiplicationOperator) SetArg2(n int) {
	op.Arg2 = n
}

func (op *MultiplicationOperator) Result() int {
	return op.Arg1 * op.Arg2
}
