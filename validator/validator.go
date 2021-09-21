// Package validator defines logic for checking if a given expression is valid
package validator

import (
	"mw_evaluator/errors"
	"mw_evaluator/errtracker"
	"mw_evaluator/models"
	"strings"
)

// IsOk ...
//
// Checks wether a passed expression is valid
func IsOk(tokens []models.Token, expression []string, endpoint string) (bool, error) {
	exprStr := strings.Join(expression, " ")

	// Edge case in case out math request doesn't start with 'What is'
	if isBeginningWhatIs(expression[0], expression[1]) {
		errtracker.Tracker.Save(exprStr, endpoint, errors.ErrNonMath)
		return false, errors.ErrNonMath
	}

	return validateRestOfExpr(tokens, endpoint, exprStr)

}

func isBeginningWhatIs(what, is string) bool {

	return what != "what" || is != "is"
}

func validateRestOfExpr(tokens []models.Token, endpoint, exprAsString string) (bool, error) {
	if len(tokens) < 2 {
		return true, nil
	}

	// Edge case of unsupported functionality
	// ex. 'What is 5 cubed?'
	if len(tokens) < 3 {
		switch tokens[1].(type) {
		case models.Operator:
			errtracker.Tracker.Save(exprAsString, endpoint, errors.ErrInvalidSyntax)
			return false, errors.ErrInvalidSyntax
		default:
			errtracker.Tracker.Save(exprAsString, endpoint, errors.ErrUnsupportedOp)
			return false, errors.ErrUnsupportedOp
		}
	}

	expectOperand := true
	for i, t := range tokens {
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
			if !expectOperand && i%2 != 0 {
				errtracker.Tracker.Save(exprAsString, endpoint, errors.ErrUnsupportedOp)
				return false, errors.ErrUnsupportedOp
			}
			errtracker.Tracker.Save(exprAsString, endpoint, errors.ErrNonMath)
			return false, errors.ErrNonMath
		}

		expectOperand = !expectOperand
	}

	return true, nil
}
