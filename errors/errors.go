// Package errors ...
package errors

import "errors"

var ErrInvalidToken = errors.New("invalid token")
var ErrUnsupportedOp = errors.New("unsupported operations")
var ErrNonMath = errors.New("non-math question")
var ErrInvalidSyntax = errors.New("invalid syntax")
