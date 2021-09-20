// Package errors contains error definitions used through our entire app
package errors

import "errors"

// ErrInvalidToken ...
//
// 	Will be set when we have an invalid token when we evaluate our expression
//	This is the only error that will not be exported to the error tracker
var ErrInvalidToken = errors.New("invalid token")

// ErrUnsupportedOp ...
//
// Will be set when we have been requested an operation we do not currently handle
var ErrUnsupportedOp = errors.New("unsupported operations")

// ErrNonMath ...
//
// Will be set when we have been requested a non-math related expression
var ErrNonMath = errors.New("non-math question")

// ErrInvalidSyntax ...
//
// Will be set when the request doesn't follow the format 'What is <expression>'
var ErrInvalidSyntax = errors.New("invalid syntax")
