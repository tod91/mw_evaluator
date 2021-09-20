// Package models defines different data structures and interfaces
// that will be used in more than 1 package
package models

import "fmt"

// OperatorID ...
//
// Simple typedef used for readability
type OperatorID string

// A Enumarator for easy filtering of the proper Operator from a token
const (
	PlusOperatorID         OperatorID = "plus"
	MinusOperatorID        OperatorID = "minus"
	DividedByOperatorID    OperatorID = "divided by"
	MultipliedByOperatorID OperatorID = "multiplied by"
)

// Token ...
//
// Interface which implements either Operator or Operand
type Token interface {
}

// Operator ...
//
// Interface which will be used to define a mathematical operator
type Operator interface {
	Evaluate(a, b int) int
}

// Operand ...
//
// Interface which will be used to define a mathematical operand
type Operand interface {
	GetValue() int
}

// Trash ...
//
// Structure that defines anything that is not an Operand or an Operator
type Trash struct{}

// Number ...
//
// Implements the Operand interface and defines a number in our expression
type Number struct {
	Value int
}

// GetValue ...
//
// Returns the value of an operand as an int
func (op *Number) GetValue() int {
	return op.Value
}

// NewOperatorFrom ...
//
// Returns a concrete Operator from a string token
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

// Plus ...
//
// Concrete Operator
type Plus struct{}

// Minus ...
//
// Concrete Operator
type Minus struct{}

// Multiply ...
//
// Concrete Operator
type Multiply struct{}

// Divide ...
//
// Concrete Operator
type Divide struct{}

// Evaluate ...
//
// Concrete implementation of the Plus interface
func (o *Plus) Evaluate(a, b int) int {
	return a + b
}

// Evaluate ...
//
// Concrete implementation of the Minus interface
func (o *Minus) Evaluate(a, b int) int {
	return a - b
}

// Evaluate ...
//
// Concrete implementation of the Multiply interface
func (o *Multiply) Evaluate(a, b int) int {
	return a * b
}

// Evaluate ...
//
// Concrete implementation of the Divide interface
func (o *Divide) Evaluate(a, b int) int {
	if b != 0 {
		return a / b
	}
	return 0
}
