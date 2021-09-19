package math

import (
	"mw_evaluator/errors"
	"mw_evaluator/models"
)

func Eval(tokens []models.Token) (int, error) {
	result := tokens[2].(models.Operand).GetValue()
	var next int
	var operator models.Operator

	expectOperand := false
	for _, t := range tokens[3:] {
		switch v := t.(type) {
		case models.Operator:
			if expectOperand {
				panic("return error")
			}

			operator = v
		case models.Operand:
			if !expectOperand {
				panic("return error")
			}

			next = v.GetValue()
			result = operator.Evaluate(result, next)

		default:
			panic(errors.ErrInvalidToken)
		}

		expectOperand = !expectOperand
	}

	return result, nil
}
