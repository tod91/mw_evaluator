package parser

import (
	"mw_evaluator/models"
	"strconv"
	"strings"
)

func isOperator(o string) bool {
	oSmall := strings.ToLower(o)

	switch oSmall {
	case "plus":
		return true
	case "minus":
		return true
	case "divided by":
		return true
	case "multiplied by":
		return true

	default:
		return false
	}
}

func operatorType(o string) models.Operator {
	oSmall := strings.ToLower(o)

	switch oSmall {
	case "plus":
		return &models.Plus{}
	case "minus":
		return &models.Minus{}
	case "divided by":
		return &models.Divide{}
	case "multiplied by":
		return &models.Multiply{}

	default:
		panic("unknown type")
	}

}

//TODO clean this up
func toTokens(w []string) []models.Token {
	var result []models.Token
	for i := 0; i < len(w); i++ {
		if w[i] == "multiplied" || w[i] == "divided" {
			if i+1 < len(w) {
				if isOperator(w[1] + " " + w[i+1]) {
					result = append(result, operatorType(w[i-1]+" "+w[i]))
				}
			}
			//if isOperator(w[1] + " " + w[i + 1]) {
			//	result = append(result, operatorType(w[i - 1] + " " + w[i]))
			//} else {
			result = append(result, models.Trash{})
			//}

		} else {
			if num, err := strconv.Atoi(w[i]); err == nil {
				result = append(result, &models.Number{num})
				continue
			}
			if isOperator(w[i]) {
				result = append(result, operatorType(w[i]))
			} else {
				result = append(result, models.Trash{})
			}
		}

	}

	return result
}

// Function for slicing out expressions into tokens
//
// one token is a word until the next white space
func Parse(expression string) []models.Token {
	return toTokens(strings.Split(expression, " "))
}
