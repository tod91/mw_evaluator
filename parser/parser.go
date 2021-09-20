// Package parser ...
package parser

import (
	"mw_evaluator/models"
	"strconv"
	"strings"
)

func PreProcessExp(expression string) []string {
	trimmedExp := toLower(expression)
	trimmedExp = strings.Trim(trimmedExp, " ")
	trimmedExp = strings.Trim(trimmedExp, "?")
	return strings.Split(trimmedExp, " ")
}

// Parse ...
// Function for slicing out expressions into tokens
//
// one token is a word until the next white space
func Parse(expression []string) []models.Token {
	return toTokens(expression)
}

func toTokens(words []string) []models.Token {
	var tokens []models.Token
	for i := 2; i < len(words); i++ {
		curr := words[i]
		if curr == "multiplied" || curr == "divided" {
			curr += " " + nextWord(words, i)
			i++
		}

		if isNumber(curr) {
			num, _ := asNumber(curr)
			tokens = append(tokens, &models.Number{Value: num})
		} else if isOperator(curr) {
			op, _ := asOperator(curr)
			tokens = append(tokens, op)
		} else {
			tokens = append(tokens, models.Trash{})
		}
	}

	return tokens
}

func isOperator(token string) bool {
	_, err := asOperator(token)
	return err == nil
}

func asOperator(token string) (models.Operator, error) {
	return models.NewOperatorFrom(token)
}

func isNumber(t string) bool {
	_, err := asNumber(t)
	return err == nil
}

func asNumber(t string) (int, error) {
	return strconv.Atoi(t)
}

func toLower(expr string) string {
	return strings.ToLower(expr)
}

func nextWord(words []string, i int) string {
	if i+1 >= len(words) {
		return ""
	}

	return words[i+1]
}
