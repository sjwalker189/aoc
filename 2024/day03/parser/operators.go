package parser

import (
	"aoc/lexer"
	"errors"
	"fmt"
)

type OperatorName int

type Operator interface {
	Name() OperatorName
	SetArg1(n int)
	SetArg2(n int)
	Result() int
}

const (
	OpMultiply = iota
	OpDont
	OpDo
)

func NewOperator(token lexer.Token) (Operator, error) {
	switch token.Kind {
	case lexer.IdentMult:
		return &ArithimiticOperator{
			name: OpMultiply,
		}, nil
	case lexer.IdentDo:
		return &ArithimiticOperator{
			name: OpDo,
		}, nil
	case lexer.IdentDont:
		return &ArithimiticOperator{
			name: OpDont,
		}, nil
	default:
		return nil, errors.New(fmt.Sprintf("unexpected lexer.TokenType: %#v", token.Kind))
	}
}

type ArithimiticOperator struct {
	name OperatorName
	Arg1 int
	Arg2 int
}

func (op *ArithimiticOperator) Name() OperatorName {
	return op.name
}

func (op *ArithimiticOperator) SetArg1(n int) {
	op.Arg1 = n
}

func (op *ArithimiticOperator) SetArg2(n int) {
	op.Arg2 = n
}

func (op *ArithimiticOperator) Result() int {
	return op.Arg1 * op.Arg2
}
