package validator

import (
	"mw_evaluator/errors"
	"mw_evaluator/errtracker"
	"mw_evaluator/models"
	"strings"
)

func IsOk(tokens []models.Token, expression []string, endpoint string) (bool, error) {
	exprStr := strings.Join(expression, " ")
	// Edge case in case out math request doesn't start with 'What is'
	if isBeginningWhatIs(expression[0], expression[1]) {
		errtracker.Tracker.Save(exprStr, endpoint, errors.ErrInvalidSyntax)
		return false, errors.ErrInvalidSyntax
	}

	return validateRestOfExpr(tokens, endpoint, exprStr)

}

func validateRestOfExpr(tokens []models.Token, endpoint, exprAsString string) (bool, error) {
	expectOperand := true
	for _, t := range tokens[2:] {
		switch t.(type) {
		case models.Operator:
			if expectOperand {
				errtracker.Tracker.Save(exprAsString, endpoint, errors.ErrInvalidSyntax)
				return false, errors.ErrInvalidSyntax
			}

		case models.Operand:
			if !expectOperand {
				errtracker.Tracker.Save(exprAsString, endpoint, errors.ErrInvalidSyntax)
				return false, errors.ErrInvalidSyntax
			}

		default:
			errtracker.Tracker.Save(exprAsString, endpoint, errors.ErrNonMath)
			return false, errors.ErrNonMath
		}

		expectOperand = !expectOperand
	}

	return true, nil
}

func isBeginningWhatIs(what, is string) bool {
	return what != "what" || is != "is"
}

// examine tokens after parse
// 1. Trash Trash Operand, operator, Operand ... -> valid
// 2. Trash trash Operand, trash, operand -> Unknown operator
// 3. trash tRASH  Operand, operator operator operand -> invalid syntax
// 4. non math expresion
