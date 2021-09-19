// Package errors ...
package errors

import "errors"

var ErrInvalidToken = errors.New("invalid token")

func errorToType(err error) string {
	switch err {
	case ErrInvalidToken:
		return "invalid token"
	default:
		return ""
	}
}
