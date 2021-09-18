package models

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
