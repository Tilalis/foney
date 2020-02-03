package compiler

import "errors"

// Errors
var (
	ErrEmptyInput           = errors.New("Empty input")
	ErrOperationNotDefined  = errors.New("Operation not defined")
	ErrUnsupportedType      = errors.New("Unsupported type")
	ErrUnsupportedOperation = errors.New("Unsupported operation")
	ErrSyntaxError          = errors.New("Syntax error")
	ErrUnexpectedEOF        = errors.New("Unexpected EOF")
	ErrDivisionByZero       = errors.New("Floating-point division by zero")
	ErrNoSuchSymbol         = errors.New("No such symbol")
	ErrEmptyStack           = errors.New("Empty stack")
)
