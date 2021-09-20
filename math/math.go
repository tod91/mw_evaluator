package math

import (
	"mw_evaluator/errors"
	"mw_evaluator/models"
)

func Eval(tokens []models.Token) int {
	// Returns number constants
	// ex. 'What is 5?'
	if len(tokens) < 2 {
		return tokens[0].(models.Operand).GetValue()
	}

	result := tokens[0].(models.Operand).GetValue()

	var next int
	var operator models.Operator

	for _, t := range tokens[1:] {
		switch v := t.(type) {
		case models.Operator:
			operator = v

		case models.Operand:
			next = v.GetValue()
			result = operator.Evaluate(result, next)

		default:
			panic(errors.ErrInvalidToken.Error())
		}
	}

	return result
}
