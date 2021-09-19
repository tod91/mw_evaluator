// Package models ...
package models

import "fmt"

type OperatorID string

const (
	PlusOperatorID         OperatorID = "plus"
	MinusOperatorID        OperatorID = "minus"
	DividedByOperatorID    OperatorID = "divided by"
	MultipliedByOperatorID OperatorID = "multiplied by"
)

// Token is an interface which implements either Operator or Operand
type Token interface {
}

type Operator interface {
	Evaluate(a, b int) int
}

type Operand interface {
	GetValue() int
}

type Trash struct{}

type Number struct {
	Value int
}

func (op *Number) GetValue() int {
	return op.Value
}

func NewOperatorFrom(token string) (Operator, error) {
	switch OperatorID(token) {
	case PlusOperatorID:
		return &Plus{}, nil
	case MinusOperatorID:
		return &Minus{}, nil
	case DividedByOperatorID:
		return &Divide{}, nil
	case MultipliedByOperatorID:
		return &Multiply{}, nil
	default:
		return nil, fmt.Errorf("can't derive operator from %v", token)
	}
}

type Plus struct{}
type Minus struct{}
type Multiply struct{}
type Divide struct{}

func (o *Plus) Evaluate(a, b int) int {
	return a + b
}

func (o *Minus) Evaluate(a, b int) int {
	return a - b
}

func (o *Multiply) Evaluate(a, b int) int {
	return a * b
}

func (o *Divide) Evaluate(a, b int) int {
	if b != 0 {
		return a / b
	}
	return 0
}
