package interpreter

import "errors"

// Errors
var (
	ErrEmptyInput            = errors.New("Empty input")
	ErrOperationNotDefined   = errors.New("Operation not defined")
	ErrUnsupportedType       = errors.New("Unsupported type")
	ErrUnsoppertedOperatrion = errors.New("Unsupported operation")
	ErrSyntaxError           = errors.New("Syntax error")
	ErrUnexpectedEOF         = errors.New("Unexpected EOF")
	ErrDivisionByZero        = errors.New("Floating-point division by zero")
	ErrNoSuchSymbol          = errors.New("No such symbol")
)
